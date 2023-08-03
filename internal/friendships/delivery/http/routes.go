package http

import (
	"go-friend-sphere/internal/friendships"

	"github.com/go-chi/chi/v5"
)

func FriendshipsRoutes(r chi.Router, h friendships.Handlers) {
	r.Post("/", h.CreateFriendship())
	r.Put("/{friendshipId}", h.UpdateFriendship())
	r.Delete("/{friendshipId}", h.DeleteFriendship())
	r.Get("/{friendshipId}", h.GetFriendshipByID())
	r.Get("/user/{userId}", h.GetFriendshipsByUserID())
}
