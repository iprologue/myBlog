package service

import (
	"encoding/json"
	"github.com/iprologue/myBlog/models"
	"github.com/iprologue/myBlog/pkg/gredis"
	"github.com/iprologue/myBlog/service/cache_service"
	"log"
)

type Article struct {
	ID            int
	TagId         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":        a.TagId,
		"title":         a.Title,
		"desc":          a.Desc,
		"content":       a.Content,
		"coverImageUrl": a.CoverImageUrl,
		"created_by":    a.CreatedBy,
		"state":         a.State,
	}

	if err := models.AddArticle(article); err != nil {
		return err
	}

	return nil
}

func (a *Article) Edit() error {
	return models.EditArticle(a.ID, map[string]interface{}{
		"tag_id":          a.TagId,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
		"modified_by":     a.ModifiedBy,
	})
}

func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			log.Println(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, article, 3600)
	return article, nil
}

func (a *Article) GetAll() ([]*models.Article, error) {
	var cacheArticles []*models.Article

	cache := cache_service.Article{
		TagId:    a.TagId,
		State:    a.State,
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetArticlesKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			log.Println(err)
		} else {
			err := json.Unmarshal(data, &cacheArticles)
			if err != nil {
				log.Println(err)
			}
			return cacheArticles, nil
		}
	}

	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, articles, 3600)
	return articles, nil
}

func (a *Article) Delete() error {
	return models.DeleteTag(a.TagId)
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticlerByID(a.ID)
}

func (a *Article) Count() (int, error) {
	return models.GetArticleToTal(a.getMaps())
}

func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}

	if a.TagId != -1 {
		maps["tag_id"] = a.TagId
	}

	return maps
}
