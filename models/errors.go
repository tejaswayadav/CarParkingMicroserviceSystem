package models

import "github.com/pkg/errors"

var (
	// ErrorModelNotFound is thrown when a model is not found
	ErrorModelNotFound = errors.New("Model Not Found")
)
