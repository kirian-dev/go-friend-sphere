package errors

import (
	"encoding/json"
	"errors" // Import the errors package from the standard library
	"net/http"
)

// Custom error messages
var ErrEmailExists = errors.New("email already exists")
var ErrFailedToHashPassword = errors.New("failed to hash password")
var ErrInvalidCredentials = errors.New("invalid credentials")

// ErrorResponse represents an error response structure.
type ErrorResponse struct {
	Message string `json:"message"`
}

// ErrorRes writes an error response to the http.ResponseWriter.
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

func NewErrorWithContext(err error, context string) error {
	return errors.New(context + ": " + err.Error())
}
