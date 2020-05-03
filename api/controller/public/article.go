package public

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/zpdev/zins/api/jsfmt"
	"github.com/zpdev/zins/common/errutils"
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
