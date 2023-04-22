package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	"mal-rest-api/internal/config"
)

type Logger struct {
	*zerolog.Logger
}

var (
	logger *Logger
	once   sync.Once
)

func GetLogger() *Logger {
	once.Do(func() {
		loggingLevel := config.GetConfig().Server.LoggingLevel
		zeroLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()

		switch loggingLevel {
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "fatal":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		case "panic":
			zerolog.SetGlobalLevel(zerolog.PanicLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.NoLevel)
		}
		logger = &Logger{&zeroLogger}
	})

	return logger
}
