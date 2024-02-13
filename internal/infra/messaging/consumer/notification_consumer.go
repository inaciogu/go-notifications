package consumer

import (
	"context"

	"github.com/inaciogu/go-notifications/internal/application/usecase"
	sqsclient "github.com/inaciogu/go-sqs/consumer"
	"github.com/inaciogu/go-sqs/consumer/message"
)

type NotificationConsumer struct {
	*sqsclient.SQSClient
}

func NewNotificationConsumer(createNotification *usecase.CreateNotification) *NotificationConsumer {
	sqsclient := sqsclient.New(nil, sqsclient.SQSClientOptions{
		QueueName: "notifications",
		Handle: func(message *message.Message) bool {
			handled := handleNotification(message, createNotification)

			return handled
		},
	})

	return &NotificationConsumer{
		sqsclient,
	}
}

func handleNotification(message *message.Message, createNotification *usecase.CreateNotification) bool {
	var input usecase.CreateNotificationInput

	err := message.Unmarshal(&input)

	if err != nil {
		return false
	}

	_, err = createNotification.Execute(context.TODO(), &input)

	return err == nil
}
