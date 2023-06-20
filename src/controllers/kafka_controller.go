package controllers

import (
	"encoding/json"
	"fmt"
	"go-crud/src/domain"
	"go-crud/src/external/providers"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaController struct{}

func NewKafkaController() *KafkaController {
	return &KafkaController{}
}

var topics = []string{
	"user-management",
}

func (KafkaController) Consume() {
	messagingProvider := providers.NewKafkaMessagingProvider()

	messageChannel := make(chan *kafka.Message)
	messagingProvider.NewConsumer(topics, messageChannel)

	for msg := range messageChannel {
		message := domain.Message{}
		err := json.Unmarshal(msg.Value, &message)

		if err != nil {
			fmt.Println(err)
		} else {
			if message.Topic == "user.created" {
				// TODO: Evoke some useCase
				fmt.Println(message.Data)
			}
		}
	}
}
