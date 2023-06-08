package main

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/shipinlei/go_getaway_demo/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	_ = lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()
	router.HttpServerRun()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
