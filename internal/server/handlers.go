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
	messagesHttp "go-friend-sphere/internal/messages/delivery/http"
	messagesRepository "go-friend-sphere/internal/messages/repository"
	messagesUsecase "go-friend-sphere/internal/messages/usecase"
	"go-friend-sphere/internal/middleware"
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
	messagesRepo := messagesRepository.NewMessageRepo(s.db)

	// useCases
	authUC := authUsecase.NewAuthUC(s.cfg, authRepo, s.logger)
	postsUC := postsUsecase.NewPostsUC(s.cfg, s.logger, postsRepo)
	commentsUC := commentsUsecase.NewCommentsUC(s.cfg, s.logger, commentsRepo)
	friendshipsUC := friendshipsUsecase.NewFriendshipsUC(s.cfg, s.logger, friendshipsRepo)
	messagesUC := messagesUsecase.NewMessagesUC(s.cfg, s.logger, messagesRepo)

	// Handlers
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUC, s.logger)
	postsHandlers := postsHttp.NewPostsHandlers(s.cfg, s.logger, postsUC)
	commentsHandlers := commentsHttp.NewCommentsHandlers(s.cfg, s.logger, commentsUC)
	friendshipsHandlers := friendshipsHttp.NewFriendshipsHandlers(s.cfg, s.logger, friendshipsUC)
	messagesHandlers := messagesHttp.NewMessagesHandlers(s.cfg, s.logger, messagesUC)

	// Create an instance of the MiddlewareManager
	middlewareManager := middleware.NewMiddlewareManager(authUC, s.cfg, s.logger)

	r.Route("/api/v1", func(r chi.Router) {
		// Register the "/auth" route group
		r.Route("/auth", func(r chi.Router) {
			authHttp.AuthRoutes(r, authHandlers, middlewareManager)
		})
		r.Route("/posts", func(r chi.Router) {
			postsHttp.PostsRoutes(r, postsHandlers, middlewareManager)
		})
		r.Route("/comments", func(r chi.Router) {
			commentsHttp.CommentsRoutes(r, commentsHandlers, middlewareManager)
		})
		r.Route("/friendships", func(r chi.Router) {
			friendshipsHttp.FriendshipsRoutes(r, friendshipsHandlers, middlewareManager)
		})
		r.Route("/messages", func(r chi.Router) {
			messagesHttp.MessagesRoutes(r, messagesHandlers, middlewareManager)
		})
	})

	http.Handle("/", r)
	return nil
}
