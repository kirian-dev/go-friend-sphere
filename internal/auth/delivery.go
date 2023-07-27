package auth

import (
	"net/http"
)

type Handlers interface {
	Register() http.HandlerFunc
}
