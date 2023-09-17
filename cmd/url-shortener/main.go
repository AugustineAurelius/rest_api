package main

import (
	"REST_API/internal/config"
	"fmt"
	"log/slog"
	"os"
)

const (
	envLocal string = "local"
	envDev   string = "dev"
	envProd  string = "prod"
)

func main() {

	cfg := config.MustLoad()

	fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	//log.With(slog.String("env", cfg.Env))

	log.Debug("debug")
	log.Info("Info")

	//TODO: init logger: slog

	//TODO: init storage: sqlite

	//TODO:init router: chi, render

	//TODO: init server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
