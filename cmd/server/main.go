package main

import (
	"fart/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	application := app.New()
	go application.GRPCSrv.MustRun()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.GRPCSrv.Stop()

}
