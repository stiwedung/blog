package controller

import (
	"net/http"
	"strings"

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

var installRedirectFilter = []string{"/install", "/static"}

func installMiddleware(ctx *gin.Context) {
	if !dao.Connected() {
		uri := ctx.Request.URL.RequestURI()
		showRedirect := true
		for _, filter := range installRedirectFilter {
			if strings.HasPrefix(uri, filter) {
				showRedirect = false
				break
			}
		}
		if showRedirect {
			ctx.Abort()
			ctx.Redirect(http.StatusSeeOther, "/install")
			return
		}
	}
	ctx.Next()
}
