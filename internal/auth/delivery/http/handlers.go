package http

import (
	"encoding/json"
	"go-friend-sphere/config"
	"go-friend-sphere/internal/auth"
	"go-friend-sphere/internal/models"
	"go-friend-sphere/pkg/errors"
	"go-friend-sphere/pkg/helpers"
	"go-friend-sphere/pkg/logger"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type authHandlers struct {
	cfg    *config.Config
	authUC auth.UseCase
	logger logger.ZapLogger
}

func NewAuthHandlers(cfg *config.Config, authUC auth.UseCase, logger logger.ZapLogger) auth.Handlers {
	return &authHandlers{cfg: cfg, authUC: authUC, logger: logger}
}

// Change the return type to http.HandlerFunc
func (h *authHandlers) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}
		if err := helpers.ReadRequest(r, user); err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}

		createdUser, err := h.authUC.Register(r.Context(), user)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusCreated, createdUser)
	}
}

func (h *authHandlers) Login() http.HandlerFunc {
	type Login struct {
		Email    string `json:"email" db:"email"`
		Password string `json:"password" db:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		login := &Login{}
		if err := helpers.ReadRequest(r, login); err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}

		user, err := h.authUC.Login(r.Context(), &models.User{
			Email:    login.Email,
			Password: login.Password,
		})
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusOK, user)
	}
}

func (h *authHandlers) GeUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usersList, err := h.authUC.GetUsers(r.Context())
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusOK, usersList)
	}
}

func (h *authHandlers) GetUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdStr := chi.URLParam(r, "userId")
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		foundedUser, err := h.authUC.GetUserById(r.Context(), userId)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusOK, foundedUser)
	}
}

func (h *authHandlers) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdStr := chi.URLParam(r, "userId")
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}

		// Parse the request body to get the JSON data from the client
		var updateUser struct {
			Email     string `json:"email"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		}

		err = json.NewDecoder(r.Body).Decode(&updateUser)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		// Create the user object with the data from the client
		user := &models.User{
			UserID:    userId,
			Email:     updateUser.Email,
			FirstName: updateUser.FirstName,
			LastName:  updateUser.LastName,
		}

		updatedUser, err := h.authUC.UpdateUser(r.Context(), user)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusOK, updatedUser)
	}
}

func (h *authHandlers) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdStr := chi.URLParam(r, "userId")
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
		}

		err = h.authUC.DeleteUser(r.Context(), userId)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusNoContent, nil)
	}
}
