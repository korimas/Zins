package public

import (
	"github.com/kataras/iris/v12"
)

type CommentRes struct {
	Ctx iris.Context
}

func (res *CommentRes) Get() (int, error) {
	// TODO: 分页
	return res.Ctx.JSON(iris.Map{"user": "get"})
}

func (res *CommentRes) Post() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "post"})
}

func (res *CommentRes) Put() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "put"})
}

type CommentDetailRes struct {
	Ctx iris.Context
}

func (res *CommentDetailRes) Get() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "get"})
}

func (res *CommentDetailRes) Post() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "post"})
}

func (res *CommentDetailRes) Put() (int, error) {
	return res.Ctx.JSON(iris.Map{"user": "put"})
}
