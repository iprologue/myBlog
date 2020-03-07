package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/iprologue/myBlog/common/errcode"
	"github.com/iprologue/myBlog/common/util"
	"github.com/iprologue/myBlog/models"
	"log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context)  {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := errcode.INVALID_PARAMS
	if ok {
		if checkAuth, _ := models.CheckAuth(username, password); checkAuth {
			token, err := util.CreateToken(username, password)
			if err != nil {
				log.Println(err.Error())
				code = errcode.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = errcode.SUCCESS
			}
		} else {
			code = errcode.ERROR_AUTH
		}
	} else  {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": errcode.GetMsg(code),
		"data": data,
	})
}