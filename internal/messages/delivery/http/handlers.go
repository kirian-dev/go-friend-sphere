package http

import (
	"encoding/json"
	"go-friend-sphere/config"
	"go-friend-sphere/internal/messages"
	"go-friend-sphere/internal/models"
	"go-friend-sphere/pkg/errors"
	"go-friend-sphere/pkg/helpers"
	"go-friend-sphere/pkg/logger"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type messagesHandlers struct {
	cfg        *config.Config
	logger     logger.ZapLogger
	messagesUC messages.UseCase
}

func NewMessagesHandlers(cfg *config.Config, logger logger.ZapLogger, messagesUC messages.UseCase) messages.Handlers {
	return &messagesHandlers{cfg: cfg, logger: logger, messagesUC: messagesUC}
}

// @Summary Create Message
// @Description create a new message
// @Tags Messages
// @Accept json
// @Produce json
// @Param message body models.Message true "Message object to be created"
// @Success 201 {object} models.Message
// @Router /messages [post]
func (h *messagesHandlers) CreateMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := &models.Message{}

		if err := helpers.ReadRequest(r, message); err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}
		if message.RecipientID == message.SenderID {
			errors.ErrorRes(w, error(http.ErrBodyNotAllowed), http.StatusBadRequest)
			return
		}

		if err := helpers.Validate(r.Context(), message); err != nil {
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		createdMessage, err := h.messagesUC.CreateMessage(r.Context(), message)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}
		helpers.WriteResponse(w, http.StatusCreated, createdMessage)
	}
}

// @Summary Update Message
// @Description update a message
// @Tags Messages
// @Param messageId path int true "Message ID"
// @Accept json
// @Produce json
// @Success 200 {object} models.Message
// @Router /messages/{messageId} [put]
func (h *messagesHandlers) UpdateMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		messageIdStr := chi.URLParam(r, "messageId")
		messageId, err := strconv.ParseInt(messageIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}
		var updateMessage struct {
			Message string `json:"message" validator:"required,gte=1,lte=1000"`
		}

		err = json.NewDecoder(r.Body).Decode(&updateMessage)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		if err := helpers.Validate(r.Context(), updateMessage); err != nil {
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		message := &models.Message{
			MessageID: messageId,
			Message:   updateMessage.Message,
		}

		updatedMessage, err := h.messagesUC.UpdateMessage(r.Context(), message)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		if updatedMessage == nil {
			helpers.WriteResponse(w, http.StatusNoContent, nil)
		} else {
			helpers.WriteResponse(w, http.StatusOK, updatedMessage)
		}
	}
}

// @Summary Delete Message
// @Description delete a message
// @Tags Messages
// @Param messageId path int true "Message ID"
// @Produce json
// @Success 204 "No Content"
// @Router /messages/{messageId} [delete]
func (h *messagesHandlers) DeleteMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		messageIdStr := chi.URLParam(r, "messageId")
		messageId, err := strconv.ParseInt(messageIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}
		err = h.messagesUC.DeleteMessage(r.Context(), messageId)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusNoContent, nil)
	}

}

// @Summary Get Message by ID
// @Description get a message by ID
// @Tags Messages
// @Param messageId path int true "Message ID"
// @Produce json
// @Success 200 {object} models.Message
// @Router /messages/{messageId} [get]
func (h *messagesHandlers) GetMessageByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		messageIdStr := chi.URLParam(r, "messageId")
		messageId, err := strconv.ParseInt(messageIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}

		foundedMessage, err := h.messagesUC.GetMessageByID(r.Context(), messageId)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusOK, foundedMessage)
	}
}

// @Summary Get Messages by User ID
// @Description get a list of messages by user ID
// @Tags Messages
// @Param userId path int true "User ID"
// @Produce json
// @Success 200 {array} models.Message
// @Router /messages/user/{userId} [get]
func (h *messagesHandlers) GetMessagesByUserID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdStr := chi.URLParam(r, "userId")
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			errors.ErrorRes(w, err, http.StatusInternalServerError)
			return
		}
		messagesList, err := h.messagesUC.GetMessagesByUserID(r.Context(), userId)
		if err != nil {
			helpers.LogError(h.logger, err)
			errors.ErrorRes(w, err, http.StatusBadRequest)
			return
		}

		helpers.WriteResponse(w, http.StatusOK, messagesList)
	}
}
