package http

import (
	"go-friend-sphere/internal/auth"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes(r chi.Router, h auth.Handlers) {
	r.Post("/register", h.Register())
}
