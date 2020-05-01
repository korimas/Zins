package controller

import (
	"github.com/kataras/iris/v12"
)

type Article struct {
	Ctx iris.Context
}

func (c *Article) Get() (int, error) {
	return c.Ctx.JSON(iris.Map{"user": "get"})
}

func (c *Article) Post() (int, error) {
	return c.Ctx.JSON(iris.Map{"user": "post"})
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
