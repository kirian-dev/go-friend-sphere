package http

import (
	"encoding/json"
	"fmt"
	"go-friend-sphere/config"
	"go-friend-sphere/internal/models"
	"go-friend-sphere/internal/posts"
	"go-friend-sphere/pkg/errors"
	"go-friend-sphere/pkg/helpers"
	"go-friend-sphere/pkg/logger"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type postsHandlers struct {
	cfg     *config.Config
	logger  logger.ZapLogger
	postsUC posts.UseCase
}

func NewPostsHandlers(cfg *config.Config, logger logger.ZapLogger, postsUC posts.UseCase) posts.Handlers {
	return &postsHandlers{cfg: cfg, logger: logger, postsUC: postsUC}
}

// @Summary Create Post
// @Description create a new post
// @Tags Posts
// @Accept json
// @Produce json
// @Param post body models.Post true "Post object to be created"
// @Success 201 {object} models.Post
// @Router /posts [post]
func (h *postsHandlers) CreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		post := &models.Post{}

		if err := helpers.ReadRequest(r, post); err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}

		if err := helpers.Validate(r.Context(), post); err != nil {
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		createdPost, err := h.postsUC.CreatePost(r.Context(), post)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}
		helpers.WriteResponse(w, http.StatusCreated, createdPost)
	}
}

// @Summary Update Post
// @Description update a post
// @Tags Posts
// @Param postId path int true "Post ID"
// @Accept json
// @Produce json
// @Success 200 {object} models.Post
// @Router /posts/{postId} [put]
func (h *postsHandlers) UpdatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postIdStr := chi.URLParam(r, "postId")
		postId, err := strconv.ParseInt(postIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}
		var updatePost struct {
			Content  string `json:"content" validator:"required,lte=2000"`
			ImageUrl string `json:"image_url" validator:"required,lte=2000"`
		}

		err = json.NewDecoder(r.Body).Decode(&updatePost)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		if err := helpers.Validate(r.Context(), updatePost); err != nil {
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		post := &models.Post{
			PostID:   postId,
			Content:  updatePost.Content,
			ImageUrl: updatePost.ImageUrl,
		}

		updatedPost, err := h.postsUC.UpdatePost(r.Context(), post)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}
		helpers.WriteResponse(w, http.StatusOK, updatedPost)
	}
}

// @Summary Delete Post
// @Description delete a post
// @Tags Posts
// @Param postId path int true "Post ID"
// @Produce json
// @Success 204 "No Content"
// @Router /posts/{postId} [delete]
func (h *postsHandlers) DeletePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postIdStr := chi.URLParam(r, "postId")
		postId, err := strconv.ParseInt(postIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}
		err = h.postsUC.DeletePost(r.Context(), postId)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusNoContent, nil)
	}

}

// @Summary Get Post by ID
// @Description get a post by ID
// @Tags Posts
// @Param postId path int true "Post ID"
// @Produce json
// @Success 200 {object} models.Post
// @Router /posts/{postId} [get]v
func (h *postsHandlers) GetPostById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postIdStr := chi.URLParam(r, "postId")
		postId, err := strconv.ParseInt(postIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}

		foundedPost, err := h.postsUC.GetPostById(r.Context(), postId)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusOK, foundedPost)
	}
}

// @Summary Get Posts
// @Description get a list of posts
// @Tags Posts
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Param search query string false "Search query"
// @Param sort query string false "Sorting order"
// @Produce json
// @Success 200 {array} models.Post
// @Router /posts [get]
func (h *postsHandlers) GetPosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		offsetStr := r.URL.Query().Get("offset")
		limitStr := r.URL.Query().Get("limit")
		query := r.URL.Query().Get("search")
		sort := r.URL.Query().Get("sort")

		offset, _ := strconv.Atoi(offsetStr)
		limit, _ := strconv.Atoi(limitStr)
		if offset < 0 {
			offset = 0
		}
		if limit <= 0 {
			limit = 10
		}

		params := models.GetPostsParams{
			Offset: offset,
			Limit:  limit,
			Query:  query,
			Sort:   sort,
		}

		postsList, err := h.postsUC.GetPosts(r.Context(), params)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusOK, postsList)
	}
}

func (h *postsHandlers) ToggleLikePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := chi.URLParam(r, "postId")
		postID, err := strconv.ParseInt(postIDStr, 10, 64)
		if err != nil {
			helpers.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid postId"})
			return
		}

		userIDStr := chi.URLParam(r, "userId")
		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			helpers.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid userId"})
			return
		}

		hasLiked, err := h.postsUC.ToggleLikePost(r.Context(), postID, userID)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}
		var msg string
		if hasLiked {
			msg = "liked"
		} else {
			msg = "not liked"
		}

		helpers.WriteResponse(w, http.StatusOK, map[string]string{"message": fmt.Sprintf("Post %s successfully", msg)})
	}
}
