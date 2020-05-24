package handlers

import (
	"github.com/tejaswayadav/CarParkingSystemMicroservice/loggers"
)

// VehicleHandler is a handler for product 'Car'
type VehicleHandler struct {
	logger *loggers.CarLogger
}

// NewVehicleHandler creates a new instance of CarHandler
func NewVehicleHandler(l *loggers.CarLogger) *VehicleHandler {
	return &VehicleHandler{l}
}
