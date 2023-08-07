package http

import (
	"go-friend-sphere/internal/comments"
	"go-friend-sphere/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func CommentsRoutes(r chi.Router, h comments.Handlers, m *middleware.MiddlewareManager) {
	r.Group(func(r chi.Router) {
		r.Use(m.JWTMiddleware)
		r.Post("/", h.CreateComment())
		r.Put("/{commentId}", h.UpdateComment())
		r.Delete("/{commentId}", h.DeleteComment())
		r.Get("/{commentId}", h.GetCommentByID())
		r.Get("/post/{postId}", h.GetCommentsByPostID())
	})

}
