package cache_service

import (
	"github.com/iprologue/myBlog/common/errcode"
	"strconv"
	"strings"
)

type Tag struct {
	ID int
	Name string
	State int

	PageNum int
	PageSize int
}

func (t *Tag) GetTagKey() string {
	keys := []string{
		errcode.CACHE_TAG,
		"LIST",
	}

	if t.Name != "" {
		keys = append(keys, t.Name)
	}
	if t.State >= 0 {
		keys = append(keys, strconv.Itoa(t.State))
	}
	if t.PageSize > 0 {
		keys = append(keys, strconv.Itoa(t.PageSize))
	}
	if t.PageNum > 0 {
		keys = append(keys, strconv.Itoa(t.PageNum))
	}

	return strings.Join(keys, "_")
}