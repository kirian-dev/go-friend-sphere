package server

import (
	authHttp "go-friend-sphere/internal/auth/delivery/http"
	authRepository "go-friend-sphere/internal/auth/repository"
	authUsecase "go-friend-sphere/internal/auth/usecase"
	postsHttp "go-friend-sphere/internal/posts/delivery/http"
	postsRepository "go-friend-sphere/internal/posts/repository"
	postsUsecase "go-friend-sphere/internal/posts/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) Handlers() error {
	// Create an instance of the Chi router
	r := chi.NewRouter()

	// Repositories
	authRepo := authRepository.NewAuthRepo(s.db)
	postsRepo := postsRepository.NewPostsRepo(s.db)

	// useCases
	authUC := authUsecase.NewAuthUC(s.cfg, authRepo, s.logger)
	postsUC := postsUsecase.NewPostsUC(s.cfg, s.logger, postsRepo)

	// Handlers
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUC, s.logger)
	postsHandlers := postsHttp.NewPostsHandlers(s.cfg, s.logger, postsUC)

	r.Route("/api/v1", func(r chi.Router) {
		// Register the "/auth" route group
		r.Route("/auth", func(r chi.Router) {
			authHttp.AuthRoutes(r, authHandlers)
		})
		r.Route("/posts", func(r chi.Router) {
			postsHttp.PostsRoutes(r, postsHandlers)
		})
	})

	http.Handle("/", r)
	return nil
}
