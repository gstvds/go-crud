package providers

import (
	"encoding/json"
	"go-crud/src/domain"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer
var name = "user-management"

type KafkaMessagingProvider struct{}

func NewKafkaMessagingProvider() *KafkaMessagingProvider {
	return &KafkaMessagingProvider{}
}

func (KafkaMessagingProvider) NewProducer() {
	var protocol string = "plaintext"
	kafkaProtocol := os.Getenv("KAFKA_PROTOCOL")
	kafkaUsername := os.Getenv("KAFKA_USERNAME")
	kafkaPassword := os.Getenv("KAFKA_PASSWORD")

	if kafkaProtocol == "SASL_SSL" {
		protocol = "sasl_ssl"
	}

	newProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"client.id":         name,
		"security.protocol": protocol,
		"sasl.username":     kafkaUsername,
		"sasl.password":     kafkaPassword,
		"sasl.mechanisms":   "plain",
	})

	if err != nil {
		panic(err)
	}

	producer = newProducer
}

func (KafkaMessagingProvider) NewConsumer(topics []string, channel chan *kafka.Message) {
	bootstrapServers := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")

	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          name,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	kafkaConsumer.SubscribeTopics(topics, nil)
	for {
		message, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			channel <- message
		}
	}
}

func (messagingProvider KafkaMessagingProvider) Send(topic string, key string, version int, message interface{}, lastCorrelationId ...string) {
	_message := domain.Message{
		Data:    message,
		Topic:   topic,
		Version: version,
	}

	byteMessage, err := json.Marshal(_message)

	if err == nil {

		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &name, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          byteMessage,
		}, nil)
	}

}
