package main

import (
	"github.com/tejaswayadav/CarParkingSystemMicroservice/handlers"
	"github.com/tejaswayadav/CarParkingSystemMicroservice/loggers"
)

func main() {
	var logger *loggers.InfoLogger
	carHandler := handlers.NewCarHandler(logger.GetLogger())
	carHandler
}
