package http

import (
	"go-friend-sphere/internal/messages"
	"go-friend-sphere/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func MessagesRoutes(r chi.Router, h messages.Handlers, m *middleware.MiddlewareManager) {
	r.Group(func(r chi.Router) {
		r.Use(m.JWTMiddleware)
		r.Post("/", h.CreateMessage())
		r.Put("/{messageId}", h.UpdateMessage())
		r.Delete("/{messageId}", h.DeleteMessage())
		r.Get("/{messageId}", h.GetMessageByID())
		r.Get("/user/{userId}", h.GetMessagesByUserID())
	})
}
