package resource

import (
	"github.com/kataras/iris/v12"
	"github.com/zpdev/zins/api"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/model"
	"github.com/zpdev/zins/product/extend"
	"github.com/zpdev/zins/service"
)

type AuthRes struct {
	Ctx iris.Context
}

func (res *AuthRes) Get() (int, error) {
	return res.Ctx.JSON(iris.Map{"auth": "get"})
}

func (res *AuthRes) Post() *api.Response {
	user := &model.User{}
	if err := res.Ctx.ReadJSON(user); err != nil {
		return api.ErrorResponse(errutils.JsonFormatError())
	}
	loginUser, err := service.AuthService.Login(extend.DB(), user)
	if err != nil {
		return api.ErrorResponse(err)
	}
	return api.NormalResponse(loginUser)
}

func (res *AuthRes) Delete() (int, error) {
	return res.Ctx.JSON(iris.Map{"auth": "delete"})
}

func (res *AuthRes) Put() (int, error) {
	return res.Ctx.JSON(iris.Map{"auth": "put"})
}
