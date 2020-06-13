package controller

import (
	"net/http"
	"strconv"

	"github.com/stiwedung/blog/config"
	"github.com/stiwedung/blog/dao"

	"github.com/gin-gonic/gin"
)

func init() {
	addController(new(installController))
}

type installController struct{}

func (ctrl *installController) relativePath() string {
	return "/install"
}

func (ctrl *installController) GET(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "install/install.html", map[string]string{"Title": "DXC博客 安装 › 配置文件"})
}

func (ctrl *installController) POST(ctx *gin.Context) {
	env := ctx.PostForm("env")
	if env == "dev" {
		config.Config.Common.ServerPort = 8080
		config.Config.Common.ReleaseMode = false
	} else {
		config.Config.Common.ServerPort = 80
		config.Config.Common.ReleaseMode = true
	}
	dbname := ctx.PostForm("dbname")
	username := ctx.PostForm("username")
	pwd := ctx.PostForm("pwd")
	dbhost := ctx.PostForm("dbhost")
	dbport := ctx.PostForm("dbport")
	config.Config.DB.DBName = dbname
	config.Config.DB.User = username
	config.Config.DB.Password = pwd
	config.Config.DB.MysqlIP = dbhost
	if port, err := strconv.Atoi(dbport); err == nil {
		config.Config.DB.MysqlPort = port
	}
	config.GenConfigFile()
	dao.Connect()
}
