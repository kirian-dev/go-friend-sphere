package posts

import (
	"net/http"
)

type Handlers interface {
	CreatePost() http.HandlerFunc
	UpdatePost() http.HandlerFunc
	DeletePost() http.HandlerFunc
	GetPost() http.HandlerFunc
	GetPosts() http.HandlerFunc
}
