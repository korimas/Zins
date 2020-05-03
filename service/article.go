package service

import (
	"github.com/jinzhu/gorm"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/model"
)

var ArticleService = &articleService{}

type articleService struct {
}

func (sec *articleService) GetArticle(db *gorm.DB, articleId uint) (*model.Article, *errutils.ZinError) {
	var article = &model.Article{}
	if db.Where("ID = ?", articleId).First(article).RecordNotFound() {
		return nil, errutils.ArticleNotFound()
	}
	return article, nil
}

func (sec *articleService) GetArticles(db *gorm.DB) ([]*model.Article, *errutils.ZinError) {
	articles := make([]*model.Article, 0)
	if err := db.Find(&articles).Error; err != nil {
		return nil, errutils.DBOperationsFailed()
	}
	return articles, nil

}

func (sec *articleService) CreateArticle(db *gorm.DB, article *model.Article) *errutils.ZinError {
	if err := db.Create(article).Error; err != nil {
		return errutils.DBOperationsFailed()
	}
	return nil
}
