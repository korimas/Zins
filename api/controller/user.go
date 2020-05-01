package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/zpdev/zins/api/schema"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/model"
	"github.com/zpdev/zins/product/extend"
	"github.com/zpdev/zins/service"
)

type User struct {
	Ctx iris.Context
}

func (c *User) Get() *schema.Response {
	users, err := service.UserService.GetUsers(extend.DB())
	if err != nil {
		return schema.ErrorResponse(err)
	}
	return schema.NormalResponse(users)
}

func (c *User) Post() *schema.Response {
	// TODO: need add on-off
	user := &model.User{}
	if err := c.Ctx.ReadJSON(user); err != nil {
		return schema.ErrorResponse(errutils.JsonFormatError())
	}
	if err := service.UserService.CreateUser(extend.DB(), user); err != nil {
		return schema.ErrorResponse(err)
	}
	user.Password = ""
	return schema.NormalResponse(user)
}

func (c *User) Delete() *schema.Response {
	var users []model.User
	if err := c.Ctx.ReadJSON(&users); err != nil {
		return schema.ErrorResponse(errutils.JsonFormatError())
	}
	if err := service.UserService.DeleteUsers(extend.DB(), users); err != nil {
		return schema.ErrorResponse(err)
	}
	return schema.NormalResponse(nil)
}

func (c *User) Put() (int, error) {
	return c.Ctx.JSON(iris.Map{"user": "put"})
}

// UserDetail
type UserDetail struct {
	Ctx iris.Context
}

func (c *UserDetail) Get() *schema.Response {
	username := c.Ctx.Params().Get("username")
	user, err := service.UserService.GetUser(extend.DB(), username)
	if err != nil {
		return schema.ErrorResponse(err)
	}
	return schema.NormalResponse(user)
}

func (c *UserDetail) Delete() *schema.Response {
	username := c.Ctx.Params().Get("username")
	if err := service.UserService.DeleteUser(extend.DB(), username); err != nil {
		return schema.ErrorResponse(err)
	}
	return schema.NormalResponse(nil)
}

func (c *UserDetail) Put() (int, error) {
	return c.Ctx.JSON(iris.Map{"user_detail": "put"})
}
