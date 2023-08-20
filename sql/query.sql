-- name: CreateNotification :one

INSERT INTO notifications (id, title, body, recipient_id, created_at, read_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: ListByRecipient :many

SELECT * FROM notifications WHERE recipient_id = $1 AND deleted_at IS NULL ORDER BY created_at DESC;