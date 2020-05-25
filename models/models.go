package models

import (
	"io"
)

// CarParkingSystemModel is an interface that incorporates all models used in this system
type CarParkingSystemModel interface {
	FromJSON(r io.Reader) error
	ToJSON(w io.Writer) error
}

// CarParkingSystemModelType is the model type of CarParkingSystem model
type CarParkingSystemModelType string

const (
	// CAR type enum
	CAR string = "car"

	// PARKING type enum
	PARKING string = "parking"
)
