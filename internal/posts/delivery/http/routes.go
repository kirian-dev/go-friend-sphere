package http

import (
	"github.com/go-chi/chi/v5"
)

func PostsRoutes(r chi.Router, h postsHandlers) {
	r.Post("/", h.CreatePost())
	r.Put("/{postId}", h.UpdatePost())
	r.Delete("/{postId}", h.DeletePost())
	r.Get("/{postId}", h.GetPost())
	r.Get("/all", h.GetPosts())

}
