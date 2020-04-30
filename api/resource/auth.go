package resource

import (
	"github.com/kataras/iris/v12"
)

type AuthRes struct {
	Ctx iris.Context
}

func (res *AuthRes) Get() (int, error) {
	return res.Ctx.JSON(iris.Map{"auth": "get"})
}

func (res *AuthRes) Post() (int, error) {
	return res.Ctx.JSON(iris.Map{"auth": "post"})
}

func (res *AuthRes) Delete() (int, error) {
	return res.Ctx.JSON(iris.Map{"auth": "delete"})
}

func (res *AuthRes) Put() (int, error) {
	return res.Ctx.JSON(iris.Map{"auth": "put"})
}
