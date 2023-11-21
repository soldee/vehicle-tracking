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

	var code = http.StatusInternalServerError

	switch {
	case errors.As(err, &invalidInputErr):
		code = http.StatusUnprocessableEntity
		if invalidInputErr.Code != 0 {
			code = invalidInputErr.Code
		}
		HandleErrorResponse(w, code, err)
	case errors.As(err, &notFoundErr):
		code = http.StatusNotFound
		if notFoundErr.Code != 0 {
			code = notFoundErr.Code
		}
		HandleErrorResponse(w, code, err)
	case errors.As(err, &internalErr):
		if internalErr.Code != 0 {
			code = internalErr.Code
		}
		HandleErrorResponse(w, code, err)
	default:
		HandleErrorResponse(w, http.StatusInternalServerError, err)
	}
}
