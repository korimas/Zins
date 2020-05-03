package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/zpdev/zins/api/jsfmt"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/model"
	"github.com/zpdev/zins/product/extend"
	"github.com/zpdev/zins/service"
	"strconv"
)

type Article struct {
	Ctx iris.Context
}

func (c *Article) Get() *jsfmt.Response {
	articles, err := service.ArticleService.GetArticles(extend.DB())
	if err != nil {
		return jsfmt.ErrorResponse(err)
	}
	return jsfmt.NormalResponse(articles)
}

func (c *Article) Post() *jsfmt.Response {
	// TODO: need add author, type ext.
	article := &model.Article{}
	if err := c.Ctx.ReadJSON(article); err != nil {
		return jsfmt.ErrorResponse(errutils.JsonFormatError(err.Error()))
	}

	if err := service.ArticleService.CreateArticle(extend.DB(), article); err != nil {
		return jsfmt.ErrorResponse(err)
	}
	return jsfmt.NormalResponse(article)
}

func (c *Article) Delete() *jsfmt.Response {
	var articles []model.Article
	if err := c.Ctx.ReadJSON(&articles); err != nil {
		return jsfmt.ErrorResponse(errutils.JsonFormatError(err.Error()))
	}
	if err := service.ArticleService.DeleteArticles(extend.DB(), articles); err != nil {
		return jsfmt.ErrorResponse(err)
	}
	return jsfmt.NormalResponse(nil)
}

func (c *Article) Put() *jsfmt.Response {
	//article := &model.Article{}
	var article map[string]interface{}
	if err := c.Ctx.ReadJSON(&article); err != nil {
		return jsfmt.ErrorResponse(errutils.JsonFormatError(err.Error()))
	}
	updatedArticle, err := service.ArticleService.UpdateArticle(extend.DB(), article)
	if err != nil {
		return jsfmt.ErrorResponse(err)
	}
	return jsfmt.NormalResponse(updatedArticle)
}

type ArticleDetail struct {
	Ctx iris.Context
}

func (c *ArticleDetail) Get() *jsfmt.Response {
	articleIdStr := c.Ctx.Params().Get("article_id")
	articleId, Cerr := strconv.Atoi(articleIdStr)
	if Cerr != nil {
		return jsfmt.ErrorResponse(errutils.ArticleNotFound())
	}

	article, err := service.ArticleService.GetArticle(extend.DB(), uint(articleId))
	if err != nil {
		return jsfmt.ErrorResponse(err)
	}

	return jsfmt.NormalResponse(article)
}

func (c *ArticleDetail) Delete() *jsfmt.Response {
	articleIdStr := c.Ctx.Params().Get("article_id")
	articleId, Cerr := strconv.Atoi(articleIdStr)
	if Cerr != nil {
		return jsfmt.ErrorResponse(errutils.ArticleNotFound())
	}
	if err := service.ArticleService.DeleteArticle(extend.DB(), uint(articleId)); err != nil {
		return jsfmt.ErrorResponse(err)
	}
	return jsfmt.NormalResponse(nil)

}

func (c *ArticleDetail) Put() (int, error) {
	return c.Ctx.JSON(iris.Map{"user": "put"})
}
