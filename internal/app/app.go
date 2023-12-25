package app

import (
	"fart/internal/app/grpcapp"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New() *App {

	grpcApp := grpcapp.New()
	return &App{
		GRPCSrv: grpcApp,
	}
}
