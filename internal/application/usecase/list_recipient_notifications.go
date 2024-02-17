package usecase

import (
	"context"

	"github.com/inaciogu/go-notifications/internal/domain/repository"
)

type ListRecipientNotifications struct {
	notificationRepository repository.NotificationRepository
}

type ListRecipientNotificationsInput struct {
	RecipientID string `json:"recipient_id"`
}

type ListRecipientNotificationsOutput struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	RecipientID string `json:"recipient_id"`
	CreatedAt   string `json:"created_at"`
	ReadAt      string `json:"read_at"`
	DeletedAt   string `json:"deleted_at"`
}

func NewListRecipientNotifications(notificationRepository repository.NotificationRepository) *ListRecipientNotifications {
	return &ListRecipientNotifications{notificationRepository: notificationRepository}
}

func (c *ListRecipientNotifications) Execute(ctx context.Context, input *ListRecipientNotificationsInput) ([]*ListRecipientNotificationsOutput, error) {
	notifications, err := c.notificationRepository.ListByRecipient(ctx, input.RecipientID)

	if err != nil {
		return nil, err
	}

	var outputNotifications []*ListRecipientNotificationsOutput

	for _, notification := range notifications {
		outputNotifications = append(outputNotifications, &ListRecipientNotificationsOutput{
			ID:          notification.ID,
			Title:       notification.Title,
			Body:        notification.Body,
			RecipientID: notification.RecipientID,
			CreatedAt:   notification.CreatedAt.String(),
			ReadAt:      notification.ReadAt.String(),
			DeletedAt:   notification.DeletedAt.String(),
		})
	}

	return outputNotifications, nil
}
