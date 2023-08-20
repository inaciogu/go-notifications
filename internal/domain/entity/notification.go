package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID          string
	Title       string
	Body        string
	RecipientID string
	CreatedAt   time.Time
	ReadAt      time.Time
	DeletedAt   time.Time
}

func NewNotification(title, body, recipientID string) (*Notification, error) {
	err := Validate(title, body, recipientID)

	if err != nil {
		return nil, err
	}

	return &Notification{
		ID:          uuid.New().String(),
		Title:       title,
		Body:        body,
		RecipientID: recipientID,
		CreatedAt:   time.Now(),
		ReadAt:      time.Time{},
		DeletedAt:   time.Time{},
	}, nil
}

func (n *Notification) Read() {
	n.ReadAt = time.Now()
}

func (n *Notification) Delete() {
	n.DeletedAt = time.Now()
}

func Validate(title, body, recipientID string) error {
	if title == "" {
		return errors.New("title is required")
	}

	if body == "" {
		return errors.New("body is required")
	}

	if recipientID == "" {
		return errors.New("recipient id is required")
	}

	return nil
}
