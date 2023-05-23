package main

import (
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/shipinlei/go_getaway_demo/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()
	router.HttpServerRun()
	fmt.Println("dev")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
