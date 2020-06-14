package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func show404Page(ctx *gin.Context) {
	ctx.String(http.StatusNotFound, "404 page")
}
