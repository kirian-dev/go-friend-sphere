package http

import (
	"go-friend-sphere/internal/auth"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes(r chi.Router, h auth.Handlers) {
	r.Post("/register", h.Register())
	r.Post("/login", h.Login())
	// r.Post("/logout", h.Logout())
	r.Get("/all", h.GeUsers())
	r.Get("/{userId}", h.GetUserById())
	r.Put("/{userId}", h.UpdateUser())
	r.Delete("/{userId}", h.DeleteUser())
}
