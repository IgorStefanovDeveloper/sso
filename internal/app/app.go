package app

import (
    "sso/internal/app/grpc"
)

type App struct {
    GRPCSrv *grpcapp.App
}

func New(
    log *slog.Logger,
    grpcPort int,
    storagePath string,
    tokenTTL time.Duration,
) *App {
    // TODO: initialized  storage

    // TODO: init auth service

    grpcApp := grpcapp.New(log, grpcPort)

    return &App {
        GRPCSrv: grpcApp,

    }
}
