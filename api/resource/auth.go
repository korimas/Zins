package resource

import (
	"github.com/kataras/iris/v12"
	"github.com/zpdev/zins/api/response"
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

func (res *AuthRes) Post() *response.Response {
	user := &model.User{}
	if err := res.Ctx.ReadJSON(user); err != nil {
		return response.ErrorResponse(errutils.JsonFormatError())
	}
	loginUser, token, err := service.AuthService.Login(extend.DB(), user)
	if err != nil {
		return response.ErrorResponse(err)
	}
	return response.NormalResponse(response.LoginResponse{
		User:  loginUser,
		Token: token,
	})
}

func (res *AuthRes) Delete() (int, error) {
	return res.Ctx.JSON(iris.Map{"auth": "delete"})
}

func (res *AuthRes) Put() (int, error) {
	return res.Ctx.JSON(iris.Map{"auth": "put"})
}
