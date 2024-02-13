package main

import (
    "sso/internal/config"
    "log/slog"
    "os"
    "sso/internal/app"
)

const (
    envLocal = "local"
    envDev   = "dev"
    envProd  = "prod"
)

func main() {
    cfg := config.MustLoad()

    log := setupLogger(cfg.Env)
//    fmt.Println(cfg)

    log.Info("starting application", slog.Any("cfg", cfg))

    application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

    application.GRPCSrv.MustRun();

    // TODO: up gRPC-service app
}



func setupLogger(env string) *slog.Logger {
    var log *slog.Logger

    switch env  {
    case envLocal:
        log = slog.New(
        slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
        )
    case envDev:
        log = slog.New(
        slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
        )
    case envProd:
         log = slog.New(
                slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
         )
    }

    return log
}
