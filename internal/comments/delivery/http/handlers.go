package http

import (
	"encoding/json"
	"go-friend-sphere/config"
	"go-friend-sphere/internal/comments"
	"go-friend-sphere/internal/models"
	"go-friend-sphere/pkg/errors"
	"go-friend-sphere/pkg/helpers"
	"go-friend-sphere/pkg/logger"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type commentsHandlers struct {
	cfg        *config.Config
	logger     logger.ZapLogger
	commentsUC comments.UseCase
}

func NewCommentsHandlers(cfg *config.Config, logger logger.ZapLogger, commentsUC comments.UseCase) comments.Handlers {
	return &commentsHandlers{cfg: cfg, logger: logger, commentsUC: commentsUC}
}

func (h *commentsHandlers) CreateComment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		comment := &models.Comment{}

		if err := helpers.ReadRequest(r, comment); err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}

		createdComment, err := h.commentsUC.CreateComment(r.Context(), comment)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}
		helpers.WriteResponse(w, http.StatusCreated, createdComment)
	}
}

func (h *commentsHandlers) UpdateComment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CommentIdStr := chi.URLParam(r, "commentId")
		CommentId, err := strconv.ParseInt(CommentIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}
		var updateComment struct {
			Message string `json:"message"`
		}

		err = json.NewDecoder(r.Body).Decode(&updateComment)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		Comment := &models.Comment{
			CommentID: CommentId,
			Message:   updateComment.Message,
		}

		updatedComment, err := h.commentsUC.UpdateComment(r.Context(), Comment)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}
		helpers.WriteResponse(w, http.StatusOK, updatedComment)
	}
}

func (h *commentsHandlers) DeleteComment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		commentIdStr := chi.URLParam(r, "commentId")
		commentId, err := strconv.ParseInt(commentIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}
		err = h.commentsUC.DeleteComment(r.Context(), commentId)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusNoContent, nil)
	}

}

func (h *commentsHandlers) GetCommentByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		commentIdStr := chi.URLParam(r, "commentId")
		commentId, err := strconv.ParseInt(commentIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}

		foundedComment, err := h.commentsUC.GetCommentByID(r.Context(), commentId)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusOK, foundedComment)
	}
}

func (h *commentsHandlers) GetCommentsByPostID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postIdStr := chi.URLParam(r, "postId")
		postId, err := strconv.ParseInt(postIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}
		commentsList, err := h.commentsUC.GetCommentsByPostID(r.Context(), postId)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusOK, commentsList)
	}
}
