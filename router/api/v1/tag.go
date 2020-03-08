package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/iprologue/myBlog/common/errcode"
	"github.com/iprologue/myBlog/common/util"
	"github.com/iprologue/myBlog/models"
	"github.com/iprologue/myBlog/pkg/setting"
	"github.com/unknwon/com"
	"net/http"
)

//获取文章标签
func GetTags(c *gin.Context) {

	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := errcode.SUCCESS

	tags, err := models.GetTags(util.GetPage(c), setting.AppSetting.PageSize, maps)
	if err != nil {
		data["list"] = tags
	}

	total, err := models.GetTagTotal(maps)
	if err != nil {
		data["tatal"] = total
	}

	c.JSON(code, gin.H{
		"code": code,
		"msg":  errcode.GetMsg(code),
		"data": data,
	})
}

// 新增文章标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := errcode.INVALID_PARAMS

	if !valid.HasErrors() {
		if !models.ExitTagByName(name) {
			code = errcode.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = errcode.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errcode.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 编辑文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}
	var state = -1
	if i := c.Query("state"); i != "" {
		state = com.StrTo(i).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := errcode.INVALID_PARAMS

	if !valid.HasErrors() {
		code = errcode.SUCCESS
		if models.ExitTagById(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)
		} else {
			code = errcode.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": errcode.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 删除文章标签
func DeletedTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := errcode.INVALID_PARAMS
	if !valid.HasErrors() {
		code = errcode.SUCCESS
		if models.ExitTagById(id) {
			models.DeleteTag(id)
		} else {
			code = errcode.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": errcode.GetMsg(code),
		"data":make(map[string]string),
	})
}
