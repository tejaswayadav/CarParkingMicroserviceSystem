package loggers

import (
	"log"
	"os"
)

// CarParkSysLogger is an interface for all logging mechanisms
type CarParkSysLogger interface {
	GetLogger() *log.Logger
}

type InfoLogger struct {
	logger *log.Logger
}

type DebugLogger struct {
	logger *log.Logger
}

type WarnLogger struct {
	logger *log.Logger
}

type ErrorLogger struct {
	logger *log.Logger
}

// GetLogger for logging of CarParkingSystem Microservice
func (info *InfoLogger) GetLogger() *InfoLogger {
	return &InfoLogger{log.New(os.Stdout, "INFO: ", log.LstdFlags)}
}
