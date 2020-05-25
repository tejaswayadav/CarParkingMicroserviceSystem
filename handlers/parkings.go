package handlers

import "github.com/tejaswayadav/CarParkingSystemMicroservice/loggers"

// ParkingHandler is a handler for intances of 'Parking'
type ParkingHandler struct {
	logger *loggers.CarParkingSystemLogger
}

// HandlerInterfaceMethod for ParkingHandler
func (p *ParkingHandler) HandlerInterfaceMethod() {

}

// NewParkingHandler creates a new instance of ParkingHandler
func NewParkingHandler(l *loggers.CarParkingSystemLogger) *ParkingHandler {
	return &ParkingHandler{l}
}
