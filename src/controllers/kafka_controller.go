package controllers

import (
	"encoding/json"
	"fmt"
	"go-crud/src/domain"
	"go-crud/src/external/providers"
	"go-crud/src/external/providers/database"
	"go-crud/src/external/repositories"
	"go-crud/src/usecases"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaController struct{}

func NewKafkaController() *KafkaController {
	return &KafkaController{}
}

var topics = []string{
	"user-management",
}

type UserCreatedMessagePayload struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type UserCreatedMessage struct {
	Data UserCreatedMessagePayload
}

func (KafkaController) Consume() {
	messagingProvider := providers.NewKafkaMessagingProvider()
	db := database.Get()
	contactRepository := repositories.NewPrismaContactRepository(db)
	createContactUseCase := usecases.NewCreateContactUseCase(contactRepository)

	messageChannel := make(chan *kafka.Message)
	go messagingProvider.Consume(topics, messageChannel)

	for msg := range messageChannel {
		message := domain.Message{}
		err := json.Unmarshal(msg.Value, &message)

		if err != nil {
			fmt.Println(err)
		} else {
			if message.Topic == "user.created" {
				parsedMessage := UserCreatedMessage{}
				if err := json.Unmarshal(msg.Value, &parsedMessage); err != nil {
					fmt.Println(err)
				} else {
					data := usecases.CreateContactInputDTO{
						UserId: parsedMessage.Data.Id,
						Email:  parsedMessage.Data.Email,
					}
					if createdContact, err := createContactUseCase.Exec(data); err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("Contact created: ", createdContact)
					}
				}
			}
		}
	}
}
