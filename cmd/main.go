package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"server/internal/common/lib/validator"
	"server/internal/config"
	"server/internal/modules/handler"
	"server/internal/modules/repository"
	"server/internal/modules/service"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	envLocal = "local"
	envProp  = "prod"
)

func main() {
	cfg := config.MustLoad()

	setupLogger(cfg.Env)

	logrus.Debug(fmt.Sprintf("Config file: %+v", cfg))

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     cfg.Database.Host,
		Username: cfg.Database.Username,
		Password: cfg.Database.Password,
		Port:     cfg.Database.Port,
		DBName:   cfg.Database.DBName,
		SSLMode:  cfg.Database.SSLMode,
	})

	if err != nil {
		logrus.Warn(fmt.Sprintf("faild to initialize db: %s", err.Error()))
		os.Exit(1)
	}

	logrus.Info(fmt.Sprintf("initialize db on port: %s", cfg.Database.Port))

	v, err := validator.SetupValidator()
	if err != nil {
		logrus.Warn(fmt.Sprintf("failed to initialize validator: %s", err.Error()))
		os.Exit(1)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(&service.Deps{
		Repos: repos,
		Cfg:   cfg,
	})
	handlers := handler.NewHandler(&handler.Deps{
		Services:  services,
		Validator: v,
	})

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      handlers.InitRoutes(),
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Error(fmt.Sprintf("failed to start server: %s", err.Error()))
		}
	}()

	logrus.Info("Startup server srarted")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

	logrus.Error("Startup server stopped successfully")
}

func setupLogger(env string) {
	switch env {
	case envLocal:
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	case envProp:
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}
