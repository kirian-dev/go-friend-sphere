package auth

import (
	"net/http"
)

type Handlers interface {
	Register() http.HandlerFunc
	Login() http.HandlerFunc
	// Logout() http.HandlerFunc
	GeUsers() http.HandlerFunc
	GetUserById() http.HandlerFunc
	UpdateUser() http.HandlerFunc
	DeleteUser() http.HandlerFunc
}
