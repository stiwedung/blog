package controller

import (
	"path/filepath"
	"reflect"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/stiwedung/blog/config"

	"github.com/gin-gonic/gin"
)

type (
	controller interface {
		relativePath() string
	}

	register interface {
		regist(*gin.RouterGroup)
	}
)

var (
	allController []controller
	allRegister   []register
	allMethod     = []string{"GET"}
)

func addController(ctrl controller) {
	allController = append(allController, ctrl)
}

func addRegister(reg register) {
	allRegister = append(allRegister, reg)
}

func MapRoutes() *gin.Engine {
	if config.Config.Common.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	route := gin.New()
	route.Use(logMiddleware, installMiddleware)
	route.Use(sessions.Sessions("user", cookie.NewStore([]byte("gnudewits235711"))))
	route.LoadHTMLGlob(filepath.Join(config.ROOT, "template/*/*.html"))
	route.Static("/static", filepath.Join(config.ROOT, "static"))

	g := route.Group("")
	rg := reflect.ValueOf(g)
	methodName := []string{"GET", "POST", "HEAD", "PUT", "DELETE", "OPTIONS"}
	methodRegister := map[string]reflect.Value{}
	for _, name := range methodName {
		methodRegister[name] = rg.MethodByName(name)
	}
	for _, ctrl := range allController {
		registController(g, ctrl, methodRegister)
	}
	for _, reg := range allRegister {
		reg.regist(g)
	}

	return route
}

func registController(g *gin.RouterGroup, ctrl controller, methodRegister map[string]reflect.Value) {
	rtype := reflect.TypeOf(ctrl)
	rval := reflect.ValueOf(ctrl)
	rPath := reflect.ValueOf(ctrl.relativePath())
	for i := 0; i < rtype.NumMethod(); i++ {
		method := rtype.Method(i)
		if method.Type.NumOut() != 0 {
			continue
		}
		if method.Type.NumIn() != 2 {
			continue
		}
		if method.Type.In(1) != reflect.TypeOf((*gin.Context)(nil)) {
			continue
		}
		register, ok := methodRegister[method.Name]
		if !ok {
			continue
		}
		register.Call([]reflect.Value{rPath, rval.MethodByName(method.Name)})
	}
}
