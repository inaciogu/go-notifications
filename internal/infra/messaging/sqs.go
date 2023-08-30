package messaging

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSClientOptions struct {
	QueueName string
}

type MessageBody struct {
	Name string `json:"name"`
}

type SQSClient struct {
	client        *sqs.SQS
	clientOptions *SQSClientOptions
}

type MessageResponse struct {
	Content string
}

func NewSQSClient(queueName string) *SQSClient {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &SQSClient{
		client: sqs.New(sess),
		clientOptions: &SQSClientOptions{
			QueueName: queueName,
		},
	}
}

func (s *SQSClient) ReceiveMessages() ([]*MessageResponse, error) {
	urlResult, err := s.client.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &s.clientOptions.QueueName,
	})

	if err != nil {
		fmt.Println(err.Error())

		return nil, err
	}

	fmt.Printf("polling messages from %s\n", s.clientOptions.QueueName)

	result, err := s.client.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:        urlResult.QueueUrl,
		WaitTimeSeconds: aws.Int64(5),
	})

	if err != nil {
		return nil, err
	}

	var messages []*MessageResponse

	for _, message := range result.Messages {
		messages = append(messages, &MessageResponse{
			Content: *message.Body,
		})

		formattedBody := strings.ReplaceAll(*message.Body, "'", "")

		var messageBody MessageBody

		err := json.Unmarshal([]byte(formattedBody), &messageBody)

		if err != nil {
			fmt.Println(err.Error())

			continue
		}

		fmt.Printf("received message: %s\n", messageBody.Name)
	}

	return messages, nil
}

func (s *SQSClient) Poll() {
	time := time.NewTicker(5 * time.Second)

	for range time.C {
		s.ReceiveMessages()
	}
}
