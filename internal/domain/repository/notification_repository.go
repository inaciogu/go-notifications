package repository

import (
	"context"

	"github.com/inaciogu/go-notifications/internal/domain/entity"
)

type NotificationRepository interface {
	Create(ctx context.Context, entity *entity.Notification) error
	ListByRecipient(ctx context.Context, recipientID string) ([]*entity.Notification, error)
}
