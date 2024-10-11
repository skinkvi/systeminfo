package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func InitLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func GetLogger() zerolog.Logger {
	return zerolog.New(os.Stdout).With().Timestamp().Logger()
}
