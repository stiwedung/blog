package controller

import (
	"net/http"
	"strings"

	"github.com/stiwedung/blog/config"
	"github.com/stiwedung/blog/dao"

	"github.com/gin-gonic/gin"
	"github.com/stiwedung/libgo/log"
)

func logMiddleware(ctx *gin.Context) {
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	ctx.Next()
	statusCode := ctx.Writer.Status()
	log.Infof("%s %s %d", method, path, statusCode)
}

func installMiddleware(ctx *gin.Context) {
	if !dao.Connected() {
		path := ctx.Request.URL.Path
		if !strings.HasPrefix(path, "/install") {
			ctx.Abort()
			ctx.Redirect(http.StatusSeeOther, "/install")
			return
		} else if config.Config.Common.ReleaseMode {
			ctx.Abort()
			ctx.Redirect(http.StatusSeeOther, "/")
			return
		}
	}
	ctx.Next()
}
