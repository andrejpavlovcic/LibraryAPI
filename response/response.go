package error

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

/*
HTTP Response handling for errors,
Returns valid JSON with error type and response code
*/

func NewErrorResponse(w http.ResponseWriter, statusCode int, response string) {
	error := ErrorResponse{
		true,
		statusCode,
		response,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&error)
}
