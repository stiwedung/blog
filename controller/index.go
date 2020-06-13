package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	addController(new(indexController))
}

type indexController struct{}

func (ctrl *indexController) relativePath() string {
	return "/"
}

func (ctrl *indexController) GET(ctx *gin.Context) {
	ctx.String(http.StatusOK, "index page")
}
