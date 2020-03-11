package app

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/iprologue/myBlog/common/errcode"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, errcode.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, errcode.ERROR
	}

	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, errcode.INVALID_PARAMS
	}

	return http.StatusOK, errcode.SUCCESS
}
