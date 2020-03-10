package cache_service

import (
	"github.com/iprologue/myBlog/common/errcode"
	"strconv"
	"strings"
)

type Article struct {
	ID int
	TagId int
	State int

	PageNum int
	PageSize int
}

func (a *Article) GetArticleKey() string {
	return errcode.CACHE_ARTICLE + "_" + strconv.Itoa(a.ID)
}

func (a *Article) GetArticlesKey() string {
	keys := []string {
		errcode.CACHE_ARTICLE,
		"LIST",
	}

	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}
	if a.TagId > 0 {
		keys = append(keys, strconv.Itoa(a.TagId))
	}
	if a.State >= 0 {
		keys = append(keys, strconv.Itoa(a.State))
	}
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}