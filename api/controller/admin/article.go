package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/zpdev/zins/api/jsfmt"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/model"
	"github.com/zpdev/zins/product/extend"
	"github.com/zpdev/zins/service"
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
	article := &model.Article{}
	if err := c.Ctx.ReadJSON(article); err != nil {
		return jsfmt.ErrorResponse(errutils.JsonFormatError())
	}

	if err := service.ArticleService.CreateArticle(extend.DB(), article); err != nil {
		return jsfmt.ErrorResponse(err)
	}
	return jsfmt.NormalResponse(article)
}

func (c *Article) Delete() (int, error) {
	return c.Ctx.JSON(iris.Map{"user": "delete"})
}

func (c *Article) Put() (int, error) {
	return c.Ctx.JSON(iris.Map{"user": "put"})
}

type ArticleDetail struct {
	Ctx iris.Context
}

func (c *ArticleDetail) Get() (int, error) {
	return c.Ctx.JSON(iris.Map{"user": "get"})
}

func (c *ArticleDetail) Post() (int, error) {
	return c.Ctx.JSON(iris.Map{"user": "post"})
}

func (c *ArticleDetail) Delete() (int, error) {
	return c.Ctx.JSON(iris.Map{"user": "delete"})
}

func (c *ArticleDetail) Put() (int, error) {
	return c.Ctx.JSON(iris.Map{"user": "put"})
}
