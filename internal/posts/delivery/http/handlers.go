package http

import (
	"go-friend-sphere/config"
	"go-friend-sphere/internal/posts"
	"go-friend-sphere/pkg/helpers"
	"go-friend-sphere/pkg/logger"
	"net/http"
)

type postsHandlers struct {
	cfg     *config.Config
	logger  logger.ZapLogger
	postsUC posts.UseCase
}

func NewPostsHandlers(cfg *config.Config, logger logger.ZapLogger, postsUC posts.UseCase) posts.Handlers {
	return &postsHandlers{cfg: cfg, logger: logger, postsUC: postsUC}
}

func (h *postsHandlers) CreatePost() http.HandlerFunc {
	return func(w *http.ResponseWriter, r *http.Request) {
		helpers.WriteResponse(w, http.StatusCreated, nil)
	}
}

func (h *postsHandlers) UpdatePost() http.HandlerFunc {
	return func(w *http.ResponseWriter, r *http.Request) {}

}

func (h *postsHandlers) DeletePost() http.HandlerFunc {
	return func(w *http.ResponseWriter, r *http.Request) {
		helpers.WriteResponse(w, http.StatusCreated, nil)
	}

}

func (h *postsHandlers) GetPost() http.HandlerFunc {
	return func(w *http.ResponseWriter, r *http.Request) {
		helpers.WriteResponse(w, http.StatusCreated, nil)
	}
}

func (h *postsHandlers) GetPosts() http.HandlerFunc {
	return func(w *http.ResponseWriter, r *http.Request) {
		helpers.WriteResponse(w, http.StatusCreated, nil)
	}
}
