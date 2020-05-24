package handlers

import "github.com/tejaswayadav/CarParkingSystemMicroservice/loggers"

// ParkingHandler is a handler for intances of 'Parking'
type ParkingHandler struct {
	logger *loggers.CarLogger
}

// NewParkingHandler creates a new instance of ParkingHandler
func NewParkingHandler(l *loggers.CarLogger) *ParkingHandler {
	return &ParkingHandler{l}
}
