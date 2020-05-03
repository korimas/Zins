package public

import (
	"github.com/kataras/iris/v12"
	"github.com/zpdev/zins/api/jsfmt"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/model"
	"github.com/zpdev/zins/product/extend"
	"github.com/zpdev/zins/service"
)

type Auth struct {
	Ctx iris.Context
}

func (c *Auth) Get() (int, error) {
	return c.Ctx.JSON(iris.Map{"auth": "get"})
}

func (c *Auth) Post() *jsfmt.Response {
	user := &model.User{}
	if err := c.Ctx.ReadJSON(user); err != nil {
		return jsfmt.ErrorResponse(errutils.JsonFormatError(err.Error()))
	}
	loginUser, token, err := service.AuthService.Login(extend.DB(), user)
	if err != nil {
		return jsfmt.ErrorResponse(err)
	}
	return jsfmt.NormalResponse(jsfmt.LoginResponse{
		User:  loginUser,
		Token: token,
	})
}

func (c *Auth) Delete() (int, error) {
	return c.Ctx.JSON(iris.Map{"auth": "delete"})
}

func (c *Auth) Put() (int, error) {
	return c.Ctx.JSON(iris.Map{"auth": "put"})
}
