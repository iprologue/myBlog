package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/iprologue/myBlog/common/errcode"
	"github.com/iprologue/myBlog/common/util"
	"github.com/iprologue/myBlog/pkg/app"
	"github.com/iprologue/myBlog/pkg/setting"
	"github.com/iprologue/myBlog/service"
	"github.com/unknwon/com"
	"net/http"
)


// @Summary 获取文章标签
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {

	appG := app.Gin{c}
	name := c.Query("name")

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	tagService := service.Tag{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	tags, err := tagService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	count, err := tagService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_COUNT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, errcode.SUCCESS, map[string]interface{}{
		"lists": tags,
		"total": count,
	})

}

type AddTagForm struct {
	Name string `form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State int `form:"state" valid:"Range(0,1)"`
}

// @Summary 新增文章标签
// @Produce  json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {

	var (
		appG = app.Gin{c}
		form AddTagForm
	)

	httpCode, errorCode := app.BindAndValid(c, &form)
	if errorCode != errcode.SUCCESS {
		appG.Response(httpCode, errorCode, nil)
		return
	}

	tagService := service.Tag{
		Name:      form.Name,
		CreatedBy: form.CreatedBy,
		State:     form.State,
	}
	exist, err := tagService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_EXIST_TAG, nil)
		return
	}

	if exist {
		appG.Response(http.StatusOK, errcode.ERROR_EXIST_TAG, nil)
		return
	}

	err = tagService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, errcode.SUCCESS,nil)

}

type EditTagForm struct {
	Id int `form:"id" valid:"Required;Min(1)"`
	Name string `form:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
	State int `form:"state" valid:"Range(0,1)"`
}

// @Summary 修改文章标签
// @Produce  json
// @Param id path int true "ID"
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param modified_by body string true "ModifiedBy"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form = EditTagForm{Id: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, code := app.BindAndValid(c, &form)
	if code != errcode.SUCCESS {
		appG.Response(httpCode, code, nil)
		return
	}

	tagService := service.Tag{
		ID:         form.Id,
		Name:       form.Name,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}

	exist, err := tagService.ExistById()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_EXIST_TAG, nil)
		return
	}

	if !exist {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = tagService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_EDIT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, errcode.SUCCESS, nil)

}

// @Summary 删除文章标签
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/{id} [delete]
func DeletedTag(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusInternalServerError, errcode.ERROR_EXIST_TAG, nil)
		return
	}

	tagService := service.Tag{ID: id}
	exist, err := tagService.ExistById()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_EXIST_TAG, nil)
		return
	}

	if !exist {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_EXIST_TAG, nil)
		return
	}

	if err := tagService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_DELETE_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, errcode.SUCCESS, nil)

}
