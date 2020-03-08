package models

import (
	"github.com/iprologue/myBlog/common/function"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func ExistArticlerByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetArticleToTal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetArticles(page int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	err := db.Preload("Tag").Where(maps).Offset(page).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil, err
	}

	return articles, nil
}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ? AND deleted_on = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	err = db.Model(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

func EditArticle(id int, data interface{}) error {
	if err := db.Model(&Article{}).Where("id = ?", id).Update(data).Error; err != nil {
		return err
	}

	return nil
}

func AddArticle(data map[string]interface{}) error {
	article := Article{
		TagID:   data["tag_id"].(int),
		Title:   data["title"].(string),
		Desc:    data["desc"].(string),
		Content: data["content"].(string),
		State:   data["state"].(int),
	}
	if err := db.Create(&article).Error; err != nil {
		return err
	}

	return nil
}

func DeleteArticle(id int) error {

	if err := db.Where("id = ?", id).Delete(&Article{}).Error; err != nil {
		return err
	}

	return nil
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", function.GetTimeUnix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", function.GetTimeUnix())

	return nil
}
