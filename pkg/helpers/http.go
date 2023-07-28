package helpers

import (
	"encoding/json"
	"go-friend-sphere/internal/models"
	"go-friend-sphere/pkg/logger"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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

func HashPassword(u *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func ComparePasswords(u *models.User, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func RemovePassword(u *models.User) {
	u.Password = ""
}
