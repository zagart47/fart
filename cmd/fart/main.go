package main

import (
	"fart/internal/app"
	"fart/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.New()

	application := app.New(cfg.Host)
	go application.GRPCSrv.MustRun()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.GRPCSrv.Stop()

}
