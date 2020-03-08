package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/iprologue/myBlog/common/errcode"
	"github.com/iprologue/myBlog/common/util"
	"github.com/iprologue/myBlog/pkg/app"
	"github.com/iprologue/myBlog/service"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context)  {
	appG := app.Gin{c}
	valid := validation.Validation{}

	username := c.Query("username")
	password := c.Query("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, errcode.INVALID_PARAMS, nil)
		return
	}

	authService := service.Auth{Username: username, Password: password}
	exist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	if !exist {
		appG.Response(http.StatusUnauthorized, errcode.ERROR_AUTH, nil)
		return
	}

	token, err := util.CreateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, errcode.SUCCESS, map[string]string{
		"token":token,
	})
}