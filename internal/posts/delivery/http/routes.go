package http

import (
	"go-friend-sphere/internal/middleware"
	"go-friend-sphere/internal/posts"

	"github.com/go-chi/chi/v5"
)

func PostsRoutes(r chi.Router, h posts.Handlers, m *middleware.MiddlewareManager) {
	r.Group(func(r chi.Router) {
		r.Use(m.JWTMiddleware)
		r.Post("/", h.CreatePost())
		r.Put("/{postId}", h.UpdatePost())
		r.Delete("/{postId}", h.DeletePost())
		r.Get("/{postId}", h.GetPostById())
		r.Get("/all", h.GetPosts())
		r.Post("/{postId}/like/{userId}", h.ToggleLikePost())
	})
}
