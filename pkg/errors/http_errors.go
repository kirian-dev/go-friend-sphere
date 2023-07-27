package errors

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents an error response structure.
type ErrorResponse struct {
	Message string `json:"message"`
}

// ErrorResponse writes an error response to the http.ResponseWriter.
func ErrorRes(w http.ResponseWriter, err error, statusCode int) {
	response := ErrorResponse{
		Message: err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		// If there's an error encoding the error response, log it or handle it as needed.
		// You can also send a simple plain text error message if JSON encoding fails.
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
