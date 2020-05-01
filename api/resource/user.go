package resource

import (
	"github.com/kataras/iris/v12"
	"github.com/zpdev/zins/api"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/model"
	"github.com/zpdev/zins/product/extend"
	"github.com/zpdev/zins/service"
)

type UserRes struct {
	Ctx iris.Context
}

func (res *UserRes) Get() *api.Response {
	users, err := service.UserService.GetUsers(extend.DB())
	if err != nil {
		return api.ErrorResponse(err)
	}
	return api.NormalResponse(users)
}

func (res *UserRes) Post() *api.Response {
	user := &model.User{}
	if err := res.Ctx.ReadJSON(user); err != nil {
		return api.ErrorResponse(errutils.JsonFormatError())
	}
	if err := service.UserService.CreateUser(extend.DB(), user); err != nil {
		return api.ErrorResponse(err)
	}
	return api.NormalResponse(user)
}

func (res *UserRes) Delete() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "delete"})
}

func (res *UserRes) Put() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "put"})
}

// UserDetail
type UserDetailRes struct {
	Ctx iris.Context
}

func (res *UserDetailRes) Get() *api.Response {
	username := res.Ctx.Params().Get("username")
	user, err := service.UserService.GetUser(extend.DB(), username)
	if err != nil {
		return api.ErrorResponse(err)
	}
	return api.NormalResponse(user)
}

func (res *UserDetailRes) Post() (int, error) {
	return res.Ctx.JSON(iris.Map{"user_detail": "post"})
}

func (res *UserDetailRes) Delete() (int, error) {
	return res.Ctx.JSON(iris.Map{"user_detail": "delete"})
}

func (res *UserDetailRes) Put() (int, error) {
	return res.Ctx.JSON(iris.Map{"user_detail": "put"})
}
