package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/skinkvi/systeminfo/internal/app"
	"github.com/skinkvi/systeminfo/internal/config"
	"github.com/skinkvi/systeminfo/pkg/logger"
)

func main() {
	logger := logger.GetLogger()

	logger.Info().Msg("Starting server...")

	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to load config")
	}

	logger.Info().Interface("config", config).Msg("Loaded config")

	app, err := app.NewApp(config, &logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create app")
	}

	logger.Info().Msg("App created")

	// Add signal handling to capture shutdown reason
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-stop
		logger.Info().Msgf("Caught signal: %v", sig)
		logger.Info().Msg("Gracefully stopping...")
		// Add cleanup code here if needed
		os.Exit(0)
	}()

	if err := app.Run(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to run app")
	}

	logger.Info().Msg("Server stopped")
}
