package helpers

import (
	"encoding/json"
	"go-friend-sphere/pkg/logger"
	"net/http"
)

func ReadRequest(r *http.Request, request interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return err
	}
	return Validate(r.Context(), request)
}
func LogError(logger logger.ZapLogger, err error) {
	if err != nil {
		logger.Error(err)
	}
}

func WriteResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
