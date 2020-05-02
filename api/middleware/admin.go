package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/zpdev/zins/api/jsfmt"
	cons "github.com/zpdev/zins/common/constance"
	"github.com/zpdev/zins/common/errutils"
	"github.com/zpdev/zins/product/cache"
	"time"
)

func AdminAuth(ctx iris.Context) {
	tokenID := ctx.GetHeader("X-User-Token")

	token := cache.TokenCache.Get(tokenID)
	if token == nil {
		failedAuth(ctx, errutils.NotLogin())
		return
	}

	if token.ExpiredAt < time.Now().Unix() {
		failedAuth(ctx, errutils.LoginTimeOut())
		return
	}

	user := cache.UserCache.Get(token.Username)
	if user == nil {
		failedAuth(ctx, errutils.UserNotFound(token.Username))
		return
	}
	if user.Role != cons.RoleAdmin {
		failedAuth(ctx, errutils.PermissionDenied())
		return
	}

	ctx.Next()

}

func failedAuth(ctx iris.Context, err *errutils.ZinError) {
	_, _ = ctx.JSON(jsfmt.ErrorResponse(err))
	ctx.StopExecution()
}
