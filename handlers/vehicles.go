package handlers

import (
	"github.com/tejaswayadav/CarParkingSystemMicroservice/loggers"
)

// VehicleHandler is a handler for product 'Car'
type VehicleHandler struct {
	logger *loggers.CarParkingSystemLogger
}

// HandlerInterfaceMethod for VehicleHandler
func (v *VehicleHandler) HandlerInterfaceMethod() {

}

// NewVehicleHandler creates a new instance of CarHandler
func NewVehicleHandler(l *loggers.CarParkingSystemLogger) *VehicleHandler {
	return &VehicleHandler{l}
}
