package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

}
