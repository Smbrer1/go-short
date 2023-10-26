package main

import (
	"log/slog"
	"os"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"

	"github.com/Smbrer1/go-short/internal/config"
	"github.com/Smbrer1/go-short/internal/helpers/logger/sl"
	"github.com/Smbrer1/go-short/internal/http-server/middleware/realip"
	"github.com/Smbrer1/go-short/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// Init Config
	cfg := config.MustLoad()

	// Init Logger
	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("Debug messages enabled")

	// Init Storage
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	// Init Router
	router := gin.New()
	// Add Middleware
	router.Use(requestid.New())
	router.Use(realip.RealIP())
	router.Use(sloggin.New(log))
	router.Use(gin.Recovery())

	// Create API Group
	v1 := router.Group("v1")
	{
		v1.GET("", nil)
	}

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
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
