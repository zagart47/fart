package grpcapp

import (
	"fart/head"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	host       string
}

func New(host string) *App {
	gRPCServer := grpc.NewServer()
	head.Register(gRPCServer)
	return &App{
		gRPCServer: gRPCServer,
		host:       host,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	l, err := net.Listen("tcp", ":11111")
	if err != nil {
		log.Fatal(err)
	}

	if err := a.gRPCServer.Serve(l); err != nil {
	}
	return nil
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
