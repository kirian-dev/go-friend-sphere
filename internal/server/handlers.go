package server

import (
	authHttp "go-friend-sphere/internal/auth/delivery/http"
	authRepository "go-friend-sphere/internal/auth/repository"
	authUsecase "go-friend-sphere/internal/auth/usecase"
	commentsHttp "go-friend-sphere/internal/comments/delivery/http"
	commentsRepository "go-friend-sphere/internal/comments/repository"
	commentsUsecase "go-friend-sphere/internal/comments/usecase"
	friendshipsHttp "go-friend-sphere/internal/friendships/delivery/http"
	friendshipsRepository "go-friend-sphere/internal/friendships/repository"
	friendshipsUsecase "go-friend-sphere/internal/friendships/usecase"
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
	commentsRepo := commentsRepository.NewCommentsRepo(s.db)
	friendshipsRepo := friendshipsRepository.NewFriendshipRepo(s.db)

	// useCases
	authUC := authUsecase.NewAuthUC(s.cfg, authRepo, s.logger)
	postsUC := postsUsecase.NewPostsUC(s.cfg, s.logger, postsRepo)
	commentsUC := commentsUsecase.NewCommentsUC(s.cfg, s.logger, commentsRepo)
	friendshipsUC := friendshipsUsecase.NewFriendshipsUC(s.cfg, s.logger, friendshipsRepo)

	// Handlers
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUC, s.logger)
	postsHandlers := postsHttp.NewPostsHandlers(s.cfg, s.logger, postsUC)
	commentsHandlers := commentsHttp.NewCommentsHandlers(s.cfg, s.logger, commentsUC)
	friendshipsHandlers := friendshipsHttp.NewFriendshipsHandlers(s.cfg, s.logger, friendshipsUC)

	r.Route("/api/v1", func(r chi.Router) {
		// Register the "/auth" route group
		r.Route("/auth", func(r chi.Router) {
			authHttp.AuthRoutes(r, authHandlers)
		})
		r.Route("/posts", func(r chi.Router) {
			postsHttp.PostsRoutes(r, postsHandlers)
		})
		r.Route("/comments", func(r chi.Router) {
			commentsHttp.CommentsRoutes(r, commentsHandlers)
		})
		r.Route("/friendships", func(r chi.Router) {
			friendshipsHttp.CommentsRoutes(r, friendshipsHandlers)
		})
	})

	http.Handle("/", r)
	return nil
}
