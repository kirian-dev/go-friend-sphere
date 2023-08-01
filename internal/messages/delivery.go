package messages

import (
	"net/http"
)

type Handlers interface {
	CreateMessage() http.HandlerFunc
	UpdateMessage() http.HandlerFunc
	DeleteMessage() http.HandlerFunc
	GetMessageByID() http.HandlerFunc
	GetMessagesByUserID() http.HandlerFunc
}
