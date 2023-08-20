package usecase

import (
	"context"

	"github.com/inaciogu/go-notifications/internal/domain/entity"
	"github.com/inaciogu/go-notifications/internal/domain/repository"
)

type CreateNotification struct {
	notificationRepository repository.NotificationRepository
}

type CreateNotificationInput struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	RecipientID string `json:"recipient_id"`
}

type CreateNotificationOutput struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	RecipientID string `json:"recipient_id"`
	CreatedAt   string `json:"created_at"`
	ReadAt      string `json:"read_at"`
	DeletedAt   string `json:"deleted_at"`
}

func NewCreateNotification(notificationRepository repository.NotificationRepository) *CreateNotification {
	return &CreateNotification{notificationRepository: notificationRepository}
}

func (c *CreateNotification) Execute(ctx context.Context, input *CreateNotificationInput) (*CreateNotificationOutput, error) {
	notification, err := entity.NewNotification(input.Title, input.Body, input.RecipientID)

	if err != nil {
		return nil, err
	}

	err = c.notificationRepository.Create(ctx, notification)

	if err != nil {
		return nil, err
	}

	return &CreateNotificationOutput{
		ID:          notification.ID,
		Title:       input.Title,
		Body:        input.Body,
		RecipientID: input.RecipientID,
		CreatedAt:   notification.CreatedAt.String(),
		ReadAt:      notification.ReadAt.String(),
		DeletedAt:   notification.DeletedAt.String(),
	}, nil
}
