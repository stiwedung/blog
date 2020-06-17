package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/stiwedung/blog/config"
	"github.com/stiwedung/blog/controller"
	"github.com/stiwedung/blog/dao"

	"github.com/stiwedung/libgo/log"
)

func init() {
	log.GetLogger().WithOption(log.WriteCallerOption(config.Config.Common.LogCaller))
}

func main() {
	dao.Connect()

	route := controller.MapRoutes()

	addr := fmt.Sprintf("0.0.0.0:%d", config.Config.Common.ServerPort)
	server := &http.Server{
		Addr:    addr,
		Handler: route,
	}

	handleSignal(server)
	log.Infof("server listening on %s", addr)
	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Errorf("server startup failed: %v", err)
		}
	}
	gracefulExit()
}

func handleSignal(server *http.Server) {
	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		c := <-s
		log.Infof("got signal %v", c)
		server.Close()
	}()
}

func gracefulExit() {
	dao.Disconnect()
	log.Info("server graceful exit")
	log.Close()
}
