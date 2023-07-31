package http

import (
	"go-friend-sphere/internal/posts"

	"github.com/go-chi/chi/v5"
)

func PostsRoutes(r chi.Router, h posts.Handlers) {
	r.Post("/", h.CreatePost())
	r.Put("/{postId}", h.UpdatePost())
	r.Delete("/{postId}", h.DeletePost())
	r.Get("/{postId}", h.GetPostById())
	r.Get("/all", h.GetPosts())
	r.Post("/{postId}/like/{userId}", h.ToggleLikePost())
}
