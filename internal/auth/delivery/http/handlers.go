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

// @Summary Register
// @Description register new user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 201 {object} models.User
// @Router /auth/register [post]
func (h *authHandlers) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}
		if err := helpers.ReadRequest(r, user); err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}

		if err := helpers.Validate(r.Context(), user); err != nil {
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		createdUser, err := h.authUC.Register(r.Context(), user)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		token, err := helpers.GenerateJWTToken(createdUser.Email, createdUser.Role, h.cfg)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"user":  createdUser,
			"token": token,
		}

		helpers.WriteResponse(w, http.StatusCreated, response)
	}
}

// @Summary Login
// @Description user login
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/login [post]
func (h *authHandlers) Login() http.HandlerFunc {
	type Login struct {
		Email    string `json:"email" db:"email" validate:"required,email,omitempty,lte=60"`
		Password string `json:"password" db:"password" validate:"required,omitempty,gte=6"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		login := &Login{}
		if err := helpers.ReadRequest(r, login); err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}

		if err := helpers.Validate(r.Context(), login); err != nil {
			errors.ErrorRes(w, err, http.StatusBadRequest)
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
		token, err := helpers.GenerateJWTToken(user.Email, user.Role, h.cfg)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}

		// Append the token to the response
		response := map[string]interface{}{
			"user":  user,
			"token": token,
		}

		helpers.WriteResponse(w, http.StatusOK, response)
	}
}

// @Summary Get Users
// @Description get a list of all users
// @Tags Auth
// @Produce json
// @Success 200 {array} models.User
// @Router /auth/users [get]
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

// @Summary Get User by ID
// @Description get a user by ID
// @Tags Auth
// @Param userId path int true "User ID"
// @Produce json
// @Success 200 {object} models.User
// @Router /auth/users/{userId} [get]
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

// @Summary Update User
// @Description update a user's details
// @Tags Auth
// @Param userId path int true "User ID"
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /auth/users/{userId} [put]
func (h *authHandlers) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdStr := chi.URLParam(r, "userId")
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}

		var updateUser struct {
			Email     string `json:"email" validate:"required,email,omitempty,lte=60"`
			FirstName string `json:"first_name" validate:"required,omitempty,lte=60"`
			LastName  string `json:"last_name" validate:"required,omitempty,lte=60"`
		}

		err = json.NewDecoder(r.Body).Decode(&updateUser)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		if err := helpers.Validate(r.Context(), updateUser); err != nil {
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

// @Summary Delete User
// @Description delete a user
// @Tags Auth
// @Param userId path int true "User ID"
// @Produce json
// @Success 204 "No Content"
// @Router /auth/users/{userId} [delete]
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
