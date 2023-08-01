package http

import (
	"encoding/json"
	"go-friend-sphere/config"
	"go-friend-sphere/internal/friendships"
	"go-friend-sphere/internal/models"
	"go-friend-sphere/pkg/errors"
	"go-friend-sphere/pkg/helpers"
	"go-friend-sphere/pkg/logger"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type friendshipsHandlers struct {
	cfg           *config.Config
	logger        logger.ZapLogger
	friendshipsUC friendships.UseCase
}

func NewFriendshipsHandlers(cfg *config.Config, logger logger.ZapLogger, friendshipsUC friendships.UseCase) friendships.Handlers {
	return &friendshipsHandlers{cfg: cfg, logger: logger, friendshipsUC: friendshipsUC}
}

func (h *friendshipsHandlers) CreateFriendship() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		friendship := &models.Friendship{}

		if err := helpers.ReadRequest(r, friendship); err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}
		if friendship.UserID == friendship.FriendID {
			errors.ErrorRes(w, error(http.ErrBodyNotAllowed), http.StatusBadRequest)
			return
		}

		createdFriendship, err := h.friendshipsUC.CreateFriendship(r.Context(), friendship)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}
		helpers.WriteResponse(w, http.StatusCreated, createdFriendship)
	}
}

func (h *friendshipsHandlers) UpdateFriendship() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		FriendshipIdStr := chi.URLParam(r, "friendshipId")
		FriendshipId, err := strconv.ParseInt(FriendshipIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}
		var updateFriendship struct {
			Status string `json:"status"`
		}

		err = json.NewDecoder(r.Body).Decode(&updateFriendship)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		friendship := &models.Friendship{
			FriendshipID: FriendshipId,
			Status:       updateFriendship.Status,
		}

		updatedFriendship, err := h.friendshipsUC.UpdateFriendship(r.Context(), friendship)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		if updatedFriendship == nil {
			helpers.WriteResponse(w, http.StatusNoContent, nil)
		} else {
			helpers.WriteResponse(w, http.StatusOK, updatedFriendship)
		}
	}
}

func (h *friendshipsHandlers) DeleteFriendship() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		friendshipIdStr := chi.URLParam(r, "friendshipId")
		friendshipId, err := strconv.ParseInt(friendshipIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}
		err = h.friendshipsUC.DeleteFriendship(r.Context(), friendshipId)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusNoContent, nil)
	}

}

func (h *friendshipsHandlers) GetFriendshipByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		friendshipIdStr := chi.URLParam(r, "friendshipId")
		friendshipId, err := strconv.ParseInt(friendshipIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}

		foundedFriendship, err := h.friendshipsUC.GetFriendshipByID(r.Context(), friendshipId)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusOK, foundedFriendship)
	}
}

func (h *friendshipsHandlers) GetFriendshipsByUserID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdStr := chi.URLParam(r, "userId")
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}
		friendshipsList, err := h.friendshipsUC.GetFriendshipsByUserID(r.Context(), userId)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusOK, friendshipsList)
	}
}
