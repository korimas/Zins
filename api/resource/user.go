package resource

import (
	"github.com/kataras/iris/v12"
	"github.com/zpdev/zins/api/response"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/model"
	"github.com/zpdev/zins/product/extend"
	"github.com/zpdev/zins/service"
)

type UserRes struct {
	Ctx iris.Context
}

func (res *UserRes) Get() *response.Response {
	users, err := service.UserService.GetUsers(extend.DB())
	if err != nil {
		return response.ErrorResponse(err)
	}
	return response.NormalResponse(users)
}

func (res *UserRes) Post() *response.Response {
	user := &model.User{}
	if err := res.Ctx.ReadJSON(user); err != nil {
		return response.ErrorResponse(errutils.JsonFormatError())
	}
	if err := service.UserService.CreateUser(extend.DB(), user); err != nil {
		return response.ErrorResponse(err)
	}
	user.Password = ""
	return response.NormalResponse(user)
}

func (res *UserRes) Delete() *response.Response {
	var users []model.User
	if err := res.Ctx.ReadJSON(&users); err != nil {
		return response.ErrorResponse(errutils.JsonFormatError())
	}
	if err := service.UserService.DeleteUsers(extend.DB(), users); err != nil {
		return response.ErrorResponse(err)
	}
	return response.NormalResponse(nil)
}

func (res *UserRes) Put() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "put"})
}

// UserDetail
type UserDetailRes struct {
	Ctx iris.Context
}

func (res *UserDetailRes) Get() *response.Response {
	username := res.Ctx.Params().Get("username")
	user, err := service.UserService.GetUser(extend.DB(), username)
	if err != nil {
		return response.ErrorResponse(err)
	}
	return response.NormalResponse(user)
}

func (res *UserDetailRes) Delete() *response.Response {
	username := res.Ctx.Params().Get("username")
	if err := service.UserService.DeleteUser(extend.DB(), username); err != nil {
		return response.ErrorResponse(err)
	}
	return response.NormalResponse(nil)
}

func (res *UserDetailRes) Put() (int, error) {
	return res.Ctx.JSON(iris.Map{"user_detail": "put"})
}
