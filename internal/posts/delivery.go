package posts

import (
	"net/http"
)

type Handlers interface {
	CreatePost() http.HandlerFunc
	UpdatePost() http.HandlerFunc
	DeletePost() http.HandlerFunc
	GetPostById() http.HandlerFunc
	GetPosts() http.HandlerFunc
	ToggleLikePost() http.HandlerFunc
}
