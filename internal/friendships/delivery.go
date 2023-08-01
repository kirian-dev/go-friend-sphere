package friendships

import (
	"net/http"
)

type Handlers interface {
	CreateFriendship() http.HandlerFunc
	UpdateFriendship() http.HandlerFunc
	DeleteFriendship() http.HandlerFunc
	GetFriendshipByID() http.HandlerFunc
	GetFriendshipsByUserID() http.HandlerFunc
}
