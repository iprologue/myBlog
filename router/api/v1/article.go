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


// 获取单个文章
func GetArticle(c *gin.Context)  {

	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, errcode.INVALID_PARAMS, nil)
		return
	}

	articleService := service.Article{ID:id}
	exist, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exist {
		appG.Response(http.StatusOK, errcode.ERROR_NOT_EXIST_ARTICLE, nil)
	}

	article, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, errcode.SUCCESS, article)

}


// 获取多个文章
func GetArticles(c *gin.Context)  {
	appG := app.Gin{c}
	valid := validation.Validation{}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state")
	}

	tagId := -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		valid.Min(tagId, 1, "tag_id")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, errcode.INVALID_PARAMS, nil)
		return
	}

	articleService := service.Article{
		TagId:    tagId,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := articleService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_COUNT_ARTICLE_FAIL, nil)
		return
	}

	articles, err := articleService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["tatal"]  = total

	appG.Response(http.StatusOK, errcode.SUCCESS, data)
}

type AddArticleForm struct {
	TagID int `form:"tag_id" valid:"Required;Min(1)"`
	Title string `form:"title" valid:"Required;MaxSize(100)"`
	Desc string `form:"desc" valid:"Required;MaxSize(255)"`
	Content string `form:"content" valid:"Required;MaxSize(65535)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State int `form:"state" valid:"Range(0,1)"`
}

// 新增文章
func AddArticle(c *gin.Context)  {
	var (
		appG = app.Gin{c}
		form AddArticleForm
	)

	httpCode, code := app.BindAndValid(c, &form)
	if code != errcode.SUCCESS {
		appG.Response(httpCode, code, nil)
		return
	}

	tagService := service.Tag{ID:form.TagID}
	exist, err := tagService.ExistById()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if !exist {
		appG.Response(http.StatusOK, errcode.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	articleService := service.Article{
		TagId:   form.TagID,
		Title:   form.Title,
		Desc:    form.Desc,
		Content: form.Content,
		State:   form.State,
	}
	if err := articleService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, errcode.SUCCESS, nil)

}

type EditArticleForm struct {
	ID int `form:"id" valid:"Required;Min(1)"`
	TagID int `form:"tag_id" valid:"Required;Min(1)"`
	Title string `form:"title" valid:"Required;MaxSize(100)"`
	Desc string `form:"desc" valid:"Required;MaxSize(255)"`
	Content string `form:"content" valid:"Required;MaxSize(65535"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(255"`
	State int `form:"state" valid:"Range(0,1)"`
}

// 修改文章
func EditArticle(c *gin.Context)  {
	var (
		appG = app.Gin{c}
		form = EditArticleForm{ID:com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, code := app.BindAndValid(c, &form)
	if code != errcode.SUCCESS {
		appG.Response(httpCode, code, nil)
		return
	}

	articleService := service.Article{
		ID:         form.ID,
		TagId:      form.TagID,
		Desc:       form.Desc,
		Content:    form.Content,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}
	exist, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if !exist {
		appG.Response(http.StatusOK, errcode.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	tagService := service.Tag{ID: form.TagID}
	exist, err = tagService.ExistById()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if !exist {
		appG.Response(http.StatusOK, errcode.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = articleService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, errcode.SUCCESS, nil)
}


// 删除文章
func DeleteArticle(c *gin.Context)  {

	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, errcode.INVALID_PARAMS, nil)
		return
	}

	articleService := service.Article{ID: id}
	exist, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exist {
		appG.Response(http.StatusOK, errcode.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	err = articleService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, errcode.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, errcode.SUCCESS, nil)

}