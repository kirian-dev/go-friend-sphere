package http

import (
	"go-friend-sphere/internal/messages"

	"github.com/go-chi/chi/v5"
)

func MessagesRoutes(r chi.Router, h messages.Handlers) {
	r.Post("/", h.CreateMessage())
	r.Put("/{messageId}", h.UpdateMessage())
	r.Delete("/{messageId}", h.DeleteMessage())
	r.Get("/{messageId}", h.GetMessageByID())
	r.Get("/user/{userId}", h.GetMessagesByUserID())
}
