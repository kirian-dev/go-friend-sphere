package server

import (
	authHttp "go-friend-sphere/internal/auth/delivery/http"
	authRepository "go-friend-sphere/internal/auth/repository"
	authUsecase "go-friend-sphere/internal/auth/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) Handlers() error {
	// Create an instance of the Chi router
	r := chi.NewRouter()

	// Repositories
	authRepo := authRepository.NewAuthRepo(s.db)

	// useCases
	authUC := authUsecase.NewAuthUC(s.cfg, authRepo, s.logger)

	// Handlers
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUC, s.logger)

	r.Route("/api/v1", func(r chi.Router) {
		// Register the "/auth" route group
		r.Route("/auth", func(r chi.Router) {
			authHttp.AuthRoutes(r, authHandlers)
		})
	})

	http.Handle("/", r)
	return nil
}
