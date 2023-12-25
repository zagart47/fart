package grpcapp

import (
	"fart/internal/service"
)

type App struct {
	GRPCSrv *service.App
}

func New() *App {

	grpcApp := service.New()
	return &App{
		GRPCSrv: grpcApp,
	}
}
