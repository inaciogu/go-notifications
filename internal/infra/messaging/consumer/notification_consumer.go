package consumer

import (
	"context"

	"github.com/inaciogu/go-notifications/internal/application/usecase"
	"github.com/inaciogu/go-sqs/consumer"
	"github.com/inaciogu/go-sqs/consumer/message"
)

type NotificationConsumer struct {
	*consumer.SQSClient
}

func NewNotificationConsumer(createNotification *usecase.CreateNotification) *NotificationConsumer {
	sqsclient := consumer.New(nil, consumer.SQSClientOptions{
		QueueName: "notifications",
		Handle: func(message *message.Message) bool {
			handled := handleNotification(message, createNotification)

			return handled
		},
		LogLevel: "debug",
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
