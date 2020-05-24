package loggers

import (
	"log"
	"os"
)

type CarLogger struct {
	InfoLogger  *log.Logger
	DebugLogger *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
}

func GetLogger() *CarLogger {
	cl := &CarLogger{}
	cl.InfoLogger = log.New(os.Stdout, "INFO: CarParkingSystem-", log.LstdFlags)
	cl.DebugLogger = log.New(os.Stdout, "DEBUG: CarParkingSystem-", log.LstdFlags)
	cl.WarnLogger = log.New(os.Stdout, "WARN: CarParkingSystem-", log.LstdFlags)
	cl.ErrorLogger = log.New(os.Stdout, "ERROR: CarParkingSystem-", log.LstdFlags)
	return cl
}
