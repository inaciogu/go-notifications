package postgres

import (
	"context"
	"database/sql"

	"github.com/inaciogu/go-notifications/internal/domain/entity"
	"github.com/inaciogu/go-notifications/internal/infra/db"
)

type PostgresNotificationRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewPostgresNotificationRepository(dbt *sql.DB) *PostgresNotificationRepository {
	return &PostgresNotificationRepository{
		DB:      dbt,
		Queries: db.New(dbt),
	}
}

func (r *PostgresNotificationRepository) Create(ctx context.Context, notification *entity.Notification) error {
	_, err := r.Queries.CreateNotification(ctx, db.CreateNotificationParams{
		ID:          notification.ID,
		Title:       notification.Title,
		Body:        notification.Body,
		RecipientID: notification.RecipientID,
		CreatedAt:   notification.CreatedAt,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresNotificationRepository) ListByRecipient(ctx context.Context, recipientID string) ([]*entity.Notification, error) {
	notifications, err := r.Queries.ListByRecipient(ctx, recipientID)

	if err != nil {
		return nil, err
	}

	var outputNotifications []*entity.Notification

	for _, notification := range notifications {
		outputNotifications = append(outputNotifications, &entity.Notification{
			ID:          notification.ID,
			Title:       notification.Title,
			Body:        notification.Body,
			RecipientID: notification.RecipientID,
			CreatedAt:   notification.CreatedAt,
			ReadAt:      notification.ReadAt.Time,
			DeletedAt:   notification.DeletedAt.Time,
		})
	}

	return outputNotifications, nil
}
