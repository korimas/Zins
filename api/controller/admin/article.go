package admin

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/zpdev/zins/api/jsfmt"
	cons "github.com/zpdev/zins/common/constance"
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
	// TODO: 增加分页,查询条件
	queryByte := []byte(c.Ctx.FormValue("query"))
	var query = jsfmt.Query{}
	if err := json.Unmarshal(queryByte, &query); err != nil {
		return jsfmt.ErrorResponse(errutils.JsonFormatError(err.Error()))
	}
	articles, err := service.ArticleService.GetArticles(extend.DB(), &query)
	if err != nil {
		return jsfmt.ErrorResponse(err)
	}
	return jsfmt.NormalResponse(articles)
}

func (c *Article) Post() *jsfmt.Response {
	article := &model.Article{}
	if err := c.Ctx.ReadJSON(article); err != nil {
		return jsfmt.ErrorResponse(errutils.JsonFormatError(err.Error()))
	}
	author := c.Ctx.Values().Get(cons.ContextUser).(*model.User)
	article.AuthorID = author.ID
	article.Author = author.Username
	article.Status = cons.ACTIVE
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
	return c.Ctx.JSON(iris.Map{"artice_detail_put": "not implement"})
}
