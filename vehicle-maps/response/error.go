package response

import (
	"log"
	"net/http"
)

func HandleErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	errorMessage := err.Error()

	log.Printf("Error code:%v, message:%s", statusCode, err)

	if statusCode >= 500 {
		errorMessage = "Service unavailable"
	}

	type ErrorResponse struct {
		Error string `json:"error"`
	}

	data := ErrorResponse{
		Error: errorMessage,
	}

	HandleJsonResponse(w, statusCode, data)
}
