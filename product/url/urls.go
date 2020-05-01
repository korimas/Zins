package url

import (
	"github.com/kataras/iris/v12/mvc"
	"github.com/zpdev/zins/api/controller"
	"github.com/zpdev/zins/product/app"
)

func Init() {
	mvc.Configure(
		app.App.Party("/rest"),
		func(m *mvc.Application) {
			m.Party("/auth").Handle(new(controller.Auth))

			m.Party("/users").Handle(new(controller.User))
			m.Party("/users/{username:string}").Handle(new(controller.UserDetail))

			m.Party("/articles").Handle(new(controller.ArticleRes))
			m.Party("/articles/{article_id:string}").Handle(new(controller.ArticleDetailRes))

			m.Party("/comments").Handle(new(controller.CommentRes))
			m.Party("/comments/{comment_id:string}").Handle(new(controller.CommentDetailRes))

			// TODO(zpzhou): admin
		})
}
