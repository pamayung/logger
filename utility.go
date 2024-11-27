package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
)

func ToField(key string, val interface{}) (field Field) {
	field = Field{
		Key: key,
		Val: val,
	}
	return
}

func WithStdout() Option {
	return func(logger *defaultLogger) error {
		// Wire STD output for both type
		logger.writers = append(logger.writers, os.Stdout)
		return nil
	}
}

func WithFileOutput(conf *OptionsFile) Option {
	return func(logger *defaultLogger) error {
		err := validator.New().Struct(conf)
		if err != nil {
			return fmt.Errorf("config for file output error: %w", err)
		}

		outputSys, err := rotateLogs.New(
			conf.FileLocation+".%Y%m%d",
			rotateLogs.WithLinkName(conf.FileLocation),
			rotateLogs.WithMaxAge(conf.FileMaxAge*24*time.Hour),
			rotateLogs.WithRotationTime(time.Hour),
		)

		if err != nil {
			return fmt.Errorf("sys file error: %w", err)
		}

		// Wire SYS config only in sys
		logger.writers = append(logger.writers, outputSys)
		logger.closer = append(logger.closer, outputSys)
		return nil
	}
}

// WithLevel set level of logger
func WithLevel(level Level) Option {
	return func(logger *defaultLogger) error {
		logger.level = level
		return nil
	}
}
