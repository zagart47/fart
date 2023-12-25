package app

import (
	"fart/internal/app/grpcapp"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(grpcPort string) *App {

	grpcApp := grpcapp.New(grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
