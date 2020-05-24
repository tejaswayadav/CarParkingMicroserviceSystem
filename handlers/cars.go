package handlers

import (
	"github.com/tejaswayadav/CarParkingSystemMicroservice/loggers"
)

// CarHandler is a handler for product 'Car'
type CarHandler struct {
	CarLogger loggers.CarParkSysLogger
}

// NewCarHandler creates a new instance of CarHandler
func NewCarHandler(l loggers.CarParkSysLogger) *CarHandler {
	return &CarHandler{l}
}
