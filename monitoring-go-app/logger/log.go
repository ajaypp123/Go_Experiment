package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger = nil

/*
LogData:
  - File: (optional) default redirect all log to stdout.
  - Hook: (optional) Provide log hook to trasfer logs.
*/
type LogData struct {
	AppName string
	Log     *log.Logger
	File    *string
	Hook    log.Hook
}

func (l *LogData) InitLogger() error {
	// Create a new logger instance
	l.Log = log.New()

	if l.File != nil {
		// Log to a file
		file, err := os.OpenFile(*l.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		l.Log.SetOutput(file)
	} else {
		// Log to stdout by default
		l.Log.SetOutput(os.Stdout)
	}

	// Use JSON formatter for structured logging
	l.Log.SetFormatter(&log.JSONFormatter{})

	// Enable caller information
	l.Log.SetReportCaller(true)

	if l.Hook != nil {
		l.Log.AddHook(l.Hook)
	}

	Logger = l.Log
	return nil
}
