package public

import (
	"github.com/kataras/iris/v12"
)

type Article struct {
	Ctx iris.Context
}

func (c *Article) Get() (int, error) {
	// TODO: 分页
	return c.Ctx.JSON(iris.Map{"user": "get"})
}

type ArticleDetail struct {
	Ctx iris.Context
}

func (c *ArticleDetail) Get() (int, error) {
	return c.Ctx.JSON(iris.Map{"user": "get"})
}
