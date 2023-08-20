package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/inaciogu/go-notifications/internal/application/usecase"
)

type NotificationHandler struct {
	CreateNotification         *usecase.CreateNotification
	ListRecipientNotifications *usecase.ListRecipientNotifications
}

func NewNotificationHandler(createNotification *usecase.CreateNotification, listRecipientNotifications *usecase.ListRecipientNotifications) *NotificationHandler {
	return &NotificationHandler{
		CreateNotification:         createNotification,
		ListRecipientNotifications: listRecipientNotifications,
	}
}

func (h *NotificationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateNotificationInput

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	notification, err := h.CreateNotification.Execute(r.Context(), &input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		println(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(notification)
}

func (h *NotificationHandler) ListManyByRecipient(w http.ResponseWriter, r *http.Request) {
	var input usecase.ListRecipientNotificationsInput

	id := chi.URLParam(r, "id")

	input.RecipientID = id

	notifications, err := h.ListRecipientNotifications.Execute(r.Context(), &input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(notifications)
}
