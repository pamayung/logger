package logger

type OptionSet struct {
	Name string `json:"name"`
	OptionsFile
}

func SetupLogger(options OptionSet) Logger {
	optionFile := OptionsFile{
		Stdout:       options.Stdout,
		FileLocation: options.FileLocation,
		FileMaxAge:   options.FileMaxAge,
		Level:        options.Level,
	}
	logging := SetupLoggerFile(options.Name, &optionFile)
	return logging
}
