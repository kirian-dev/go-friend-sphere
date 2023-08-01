package http

import (
	"go-friend-sphere/internal/comments"

	"github.com/go-chi/chi/v5"
)

func CommentsRoutes(r chi.Router, h comments.Handlers) {
	r.Post("/", h.CreateComment())
	r.Put("/{commentId}", h.UpdateComment())
	r.Delete("/{commentId}", h.DeleteComment())
	r.Get("/{commentId}", h.GetCommentByID())
	r.Get("/post/{postId}", h.GetCommentsByPostID())
}
