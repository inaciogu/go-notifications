// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createNotification = `-- name: CreateNotification :one

INSERT INTO notifications (id, title, body, recipient_id, created_at, read_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, recipient_id, body, title, created_at, read_at, deleted_at
`

type CreateNotificationParams struct {
	ID          string
	Title       string
	Body        string
	RecipientID string
	CreatedAt   time.Time
	ReadAt      sql.NullTime
	DeletedAt   sql.NullTime
}

func (q *Queries) CreateNotification(ctx context.Context, arg CreateNotificationParams) (Notification, error) {
	row := q.db.QueryRowContext(ctx, createNotification,
		arg.ID,
		arg.Title,
		arg.Body,
		arg.RecipientID,
		arg.CreatedAt,
		arg.ReadAt,
		arg.DeletedAt,
	)
	var i Notification
	err := row.Scan(
		&i.ID,
		&i.RecipientID,
		&i.Body,
		&i.Title,
		&i.CreatedAt,
		&i.ReadAt,
		&i.DeletedAt,
	)
	return i, err
}

const listByRecipient = `-- name: ListByRecipient :many

SELECT id, recipient_id, body, title, created_at, read_at, deleted_at FROM notifications WHERE recipient_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC
`

func (q *Queries) ListByRecipient(ctx context.Context, recipientID string) ([]Notification, error) {
	rows, err := q.db.QueryContext(ctx, listByRecipient, recipientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Notification
	for rows.Next() {
		var i Notification
		if err := rows.Scan(
			&i.ID,
			&i.RecipientID,
			&i.Body,
			&i.Title,
			&i.CreatedAt,
			&i.ReadAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
