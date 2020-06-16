package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stiwedung/blog/service"
)

func init() {
	ctrl := new(articleController)
	addController(ctrl)
	addRegister(ctrl)
}

type articleController struct{}

func (ctrl *articleController) relativePath() string {
	return "/article/:id"
}

func (ctrl *articleController) regist(g *gin.RouterGroup) {
	g.GET("/list/article", articleList)
	g.GET("/editor", articleEditor)
	g.POST("/write", articleWrite)
}

func (ctrl *articleController) GET(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		show404Page(ctx)
		return
	}
	article := service.ShowArticle(id)
	ctx.String(http.StatusOK, "show article %d %s\n%s", article.ID, article.Title, article.Content)
}

func articleList(ctx *gin.Context) {
	lst := service.ArticleList()
	if len(lst) == 0 {
		ctx.String(http.StatusOK, "no article")
		return
	}
	ctx.String(http.StatusOK, "article %s", lst[0].Title)
}

func articleEditor(ctx *gin.Context) {

}

func articleWrite(ctx *gin.Context) {

}
