package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stiwedung/blog/service"
)

func init() {
	addController(new(adminController))
}

type adminController struct{}

func (ctrl *adminController) relativePath() string {
	return "/admin"
}

func (ctrl *adminController) GET(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/admin.html", map[string]string{
		"Title": "DXC博客 › 管理员",
	})
}

func (ctrl *adminController) POST(ctx *gin.Context) {
	var isLocal bool
	if ctx.Request.RemoteAddr == "127.0.0.1" {
		isLocal = true
	}
	username := ctx.PostForm("username")
	passwd := ctx.PostForm("passwd")
	err := service.Login(username, passwd, isLocal)
	if err != nil {
		show404Page(ctx)
		return
	}
	session := sessions.Default(ctx)
	session.Set(userInfo, &sessionData{
		UserName: username,
	})
	ctx.Redirect(http.StatusPermanentRedirect, "/")
	session.Save()
}
