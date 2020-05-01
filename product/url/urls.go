package url

import (
	"github.com/kataras/iris/v12/mvc"
	"github.com/zpdev/zins/api/controller"
	"github.com/zpdev/zins/product/app"
)

func Init() {
	mvc.Configure(app.App.Party("/rest/v1"), configureV1)
	mvc.Configure(app.App.Party("/rest/v1/admin"), configureV1Admin)
}

func configureV1(m *mvc.Application) {
	m.Party("/auth").Handle(&controller.Auth{}).Router.Use()

	m.Party("/users").Handle(&controller.User{})
	m.Party("/users/{username:string}").Handle(&controller.UserDetail{})

	m.Party("/articles").Handle(&controller.Article{})
	m.Party("/articles/{article_id:string}").Handle(&controller.ArticleDetail{})

	m.Party("/comments").Handle(new(controller.CommentRes))
	m.Party("/comments/{comment_id:string}").Handle(new(controller.CommentDetailRes))
}

func configureV1Admin(m *mvc.Application) {
	m.Router.Use(middleware.AdminAuth)
	m.Party("/users").Handle(&controller.Auth{}).Router.Use()
}
