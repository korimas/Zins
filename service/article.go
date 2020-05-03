package service

import (
	set "github.com/deckarep/golang-set"
	"github.com/jinzhu/gorm"
	"github.com/zpdev/zins/api/jsfmt"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/model"
	"time"
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

func (sec *articleService) GetArticles(db *gorm.DB, query *jsfmt.Query) ([]*model.Article, *errutils.ZinError) {
	var Zerr *errutils.ZinError
	articles := make([]*model.Article, 0)
	db, Zerr = query.GenDB(db)
	if Zerr != nil {
		return nil, Zerr
	}
	if err := db.Find(&articles).Error; err != nil {
		return nil, errutils.DBOperationsFailed(err.Error())
	}
	return articles, nil

}

func (sec *articleService) CreateArticle(db *gorm.DB, article *model.Article) *errutils.ZinError {
	article.CreatedTime = time.Now().Unix()
	if err := db.Create(article).Error; err != nil {
		return errutils.DBOperationsFailed(err.Error())
	}
	return nil
}

func (sec *articleService) DeleteArticle(db *gorm.DB, articleId uint) *errutils.ZinError {
	if err := db.Where("ID = ?", articleId).Delete(&model.Article{}).Error; err != nil {
		return errutils.DBOperationsFailed(err.Error())
	}
	return nil
}

func (sec *articleService) DeleteArticles(db *gorm.DB, articles []model.Article) *errutils.ZinError {
	needDeleteArticles := set.NewSet()
	for i := 0; i < len(articles); i++ {
		needDeleteArticles.Add(articles[i].ID)
	}
	if err := db.Where("ID in (?)", needDeleteArticles.ToSlice()).Delete(&model.Article{}).Error; err != nil {
		return errutils.DBOperationsFailed(err.Error())
	}
	return nil
}

func (sec *articleService) UpdateArticle(db *gorm.DB, article map[string]interface{}) (*model.Article, *errutils.ZinError) {
	updateArticle := &model.Article{}
	if db.Where("ID = ?", article["id"]).First(updateArticle).RecordNotFound() {
		return nil, errutils.ArticleNotFound()
	}

	article["updated_time"] = time.Now().Unix()
	if err := db.Model(updateArticle).Update(article).Error; err != nil {
		return nil, errutils.DBOperationsFailed(err.Error())
	}

	db.Where("ID = ?", article["id"]).First(updateArticle)
	return updateArticle, nil
}
