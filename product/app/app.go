package app

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

var App *iris.Application

func Init() {
	App = iris.New()
	App.Logger().SetLevel("debug")
	App.Use(recover.New())
	App.Use(logger.New())
}

func Run() {
	App.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
