package http

import (
	"go-friend-sphere/config"
	"go-friend-sphere/internal/auth"
	"go-friend-sphere/internal/models"
	"go-friend-sphere/pkg/errors"
	"go-friend-sphere/pkg/helpers"
	"go-friend-sphere/pkg/logger"
	"net/http"
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
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		createdUser, err := h.authUC.Register(r.Context(), user)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}

		helpers.WriteResponse(w, http.StatusCreated, createdUser)
	}
}
