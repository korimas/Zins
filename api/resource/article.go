package resource

import (
	"github.com/kataras/iris/v12"
)

type ArticleRes struct {
	Ctx iris.Context
}

func (res *ArticleRes) Get() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "get"})
}

func (res *ArticleRes) Post() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "post"})
}

func (res *ArticleRes) Delete() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "delete"})
}

func (res *ArticleRes) Put() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "put"})
}

type ArticleDetailRes struct {
	Ctx iris.Context
}

func (res *ArticleDetailRes) Get() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "get"})
}

func (res *ArticleDetailRes) Post() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "post"})
}

func (res *ArticleDetailRes) Delete() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "delete"})
}

func (res *ArticleDetailRes) Put() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "put"})
}
