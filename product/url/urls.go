package url

import (
	"github.com/kataras/iris/v12/mvc"
	"github.com/zpdev/zins/api/resource"
	"github.com/zpdev/zins/product/app"
)

func Init() {
	mvc.Configure(
		app.App.Party("/rest"),
		func(m *mvc.Application) {
			m.Party("/auth").Handle(new(resource.AuthRes))

			m.Party("/users").Handle(new(resource.UserRes))
			m.Party("/users/{username:string}").Handle(new(resource.UserDetailRes))

			m.Party("/articles").Handle(new(resource.ArticleRes))
			m.Party("/articles/{article_id:string}").Handle(new(resource.ArticleDetailRes))

			m.Party("/comments").Handle(new(resource.CommentRes))
			m.Party("/comments/{comment_id:string}").Handle(new(resource.CommentDetailRes))
		})
}
