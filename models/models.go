package models

import (
	"fmt"
	"github.com/iprologue/myBlog/common/function"
	"github.com/iprologue/myBlog/pkg/setting"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {

	var err error
	db, err = gorm.Open(setting.DataBaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DataBaseSetting.User,
		setting.DataBaseSetting.Password,
		setting.DataBaseSetting.Host,
		setting.DataBaseSetting.Name))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DataBaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		unix := function.GetTimeUnix()
		if field, ok := scope.FieldByName("CreatedOn"); ok {
			if field.IsBlank {
				err := field.Set(unix)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}

		if field, ok := scope.FieldByName("ModifiedOn"); ok {
			if field.IsBlank {
				err := field.Set(unix)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_collum"); !ok {
		err := scope.SetColumn("Modified0n", function.GetTimeUnix())
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func CloseDB() {
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalln("close DB err: ", err)
		}
	}()

}
