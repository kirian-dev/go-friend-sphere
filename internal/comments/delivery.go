package comments

import (
	"net/http"
)

type Handlers interface {
	CreateComment() http.HandlerFunc
	UpdateComment() http.HandlerFunc
	DeleteComment() http.HandlerFunc
	GetCommentByID() http.HandlerFunc
	GetCommentsByPostID() http.HandlerFunc
}
