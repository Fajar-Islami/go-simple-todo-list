package helper

import (
	"errors"

	"github.com/rs/zerolog/log"
)

const (
	LoggerLevelTrace = "LoggerLevelTrace"
	LoggerLevelDebug = "LoggerLevelDebug"
	LoggerLevelInfo  = "LoggerLevelInfo"
	LoggerLevelWarn  = "LoggerLeveWarn"
	LoggerLevelError = "LoggerLevelError"
	LoggerLevelFatal = "LoggerLevelFatal"
	LoggerLevelPanic = "LoggerLevelPanic"
)

func Logger(filepath, level, message string, err error) {
	if err == nil && (filepath == "" || level == "" || message == "") {
		log.Error().Stack().Err(errors.New("all params log is required")).Msg("")
	}

	switch level {
	case LoggerLevelDebug:
		log.Debug().Msg(message)
	case LoggerLevelInfo:
		log.Info().Msg(message)
	case LoggerLevelWarn:
		log.Warn().Msg(message)
	case LoggerLevelError:
		log.Error().Err(err)
	case LoggerLevelFatal:
		log.Fatal().Err(err)
	case LoggerLevelPanic:
		log.Panic().Err(err)
	default:
		log.Error().Err(errors.New("logger level invalid"))
	}

}
