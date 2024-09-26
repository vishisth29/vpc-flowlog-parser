package logger

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
)

// InitLogger initializes the loggers
func InitLogger() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Infof logs informational messages
func Infof(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
}

// Warnf logs warning messages
func Warnf(format string, v ...interface{}) {
	warnLogger.Printf(format, v...)
}

// Errorf logs error messages
func Errorf(format string, v ...interface{}) {
	errorLogger.Printf(format, v...)
}
