package http

import (
	"go-friend-sphere/internal/auth"
	"go-friend-sphere/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes(r chi.Router, h auth.Handlers, m *middleware.MiddlewareManager) {
	r.Post("/register", h.Register())
	r.Post("/login", h.Login())
	// r.Post("/logout", h.Logout())

	r.Group(func(r chi.Router) {
		r.Use(m.JWTMiddleware)

		r.Get("/all", h.GeUsers())
		r.Get("/{userId}", h.GetUserById())
		r.Put("/{userId}", h.UpdateUser())
		r.Delete("/{userId}", h.DeleteUser())
	})
}
