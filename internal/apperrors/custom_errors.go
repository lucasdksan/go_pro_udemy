package apperrors

import (
	"errors"
	"net/http"
)

var ErrorBadRequest = func(text string) error {
	return NewWithStatus(errors.New(text), http.StatusBadRequest)
}

var ErrorInternalServer = func(text string) error {
	return NewWithStatus(errors.New(text), http.StatusInternalServerError)
}

var ErrorNotFound = func(text string) error {
	return NewWithStatus(errors.New(text), http.StatusNotFound)
}
