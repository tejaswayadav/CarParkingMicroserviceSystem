package databases

import "github.com/pkg/errors"

var (
	// ErrorDatabaseNotFound to be called when no database is found by the provided database name
	ErrorDatabaseNotFound = errors.New("Database Not Found")
)
