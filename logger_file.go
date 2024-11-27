package logger

import (
	"fmt"
	"time"
)

type OptionsFile struct {
	Stdout       bool          `json:"stdout"`
	FileLocation string        `json:"fileLocation"`
	FileMaxAge   time.Duration `json:"fileMaxAge"`
	Level        Level         `json:"level"`
}

// SetupLoggerFile will return legacy Logger using File interface with new logic using Logger
func SetupLoggerFile(serviceName string, config *OptionsFile) Logger {
	fmt.Println("Checking newLogger File...")

	if config == nil {
		panic("legacy logger file config is nil")
	}

	var opt = make([]Option, 0)

	if config.Stdout {
		opt = append(opt, WithStdout())
	} else {
		opt = append(opt, WithFileOutput(config))
	}

	opt = append(opt, WithLevel(config.Level))

	log, err := newLogger(opt...)
	if err != nil {
		panic(fmt.Errorf("init legacy logger with mode %s error: %w", File, err))
	}

	return log
}
