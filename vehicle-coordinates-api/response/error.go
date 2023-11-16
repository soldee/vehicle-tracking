package response

import (
	"errors"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func HandleErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	errorMessage := err.Error()

	if statusCode >= 500 {
		log.Printf("error code %v with message: %s", statusCode, err)
		errorMessage = "Service unavailable"
	}

	data := ErrorResponse{
		Error: errorMessage,
	}

	HandleJsonResponse(w, statusCode, data)
}

func HandleErrorResponseErr(w http.ResponseWriter, err error) {

	var internalErr *InternalError
	var invalidInputErr *InvalidInput
	var notFoundErr *NotFound

	switch {
	case errors.As(err, &invalidInputErr):
		HandleErrorResponse(w, http.StatusUnprocessableEntity, err)
	case errors.As(err, &notFoundErr):
		HandleErrorResponse(w, http.StatusNotFound, err)
	case errors.As(err, &internalErr):
		HandleErrorResponse(w, http.StatusInternalServerError, err)
	default:
		HandleErrorResponse(w, http.StatusInternalServerError, err)
	}
}
