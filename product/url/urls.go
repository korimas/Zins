package url

import (
	"github.com/kataras/iris/v12/mvc"
	"github.com/zpdev/zins/api/controller/admin"
	"github.com/zpdev/zins/api/middleware"
	"github.com/zpdev/zins/product/app"
)

func Init() {
	mvc.Configure(app.App.Party("/rest/v1"), configureV1)
	mvc.Configure(app.App.Party("/rest/v1/admin"), configureV1Admin)
}

func configureV1(m *mvc.Application) {
	m.Party("/auth").Handle(&admin.Auth{})

	m.Party("/users").Handle(&admin.User{})
	m.Party("/users/{username:string}").Handle(&admin.UserDetail{})

	m.Party("/articles").Handle(&admin.Article{})
	m.Party("/articles/{article_id:string}").Handle(&admin.ArticleDetail{})

	m.Party("/comments").Handle(new(admin.CommentRes))
	m.Party("/comments/{comment_id:string}").Handle(new(admin.CommentDetailRes))
}

func configureV1Admin(m *mvc.Application) {
	m.Router.Use(middleware.AdminAuth)

	m.Party("/users").Handle(&admin.User{})
	m.Party("/users/{username:string}").Handle(&admin.UserDetail{})

	m.Party("/articles").Handle(&admin.Article{})
	m.Party("/articles/{article_id:string}").Handle(&admin.ArticleDetail{})

	m.Party("/comments").Handle(new(admin.CommentRes))
	m.Party("/comments/{comment_id:string}").Handle(new(admin.CommentDetailRes))
}
