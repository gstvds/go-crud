package providers

import (
	"encoding/json"
	"go-crud/src/domain"
	"go-crud/src/shared/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer
var name = "user-management"

type KafkaMessagingProvider struct{}

func NewKafkaMessagingProvider() *KafkaMessagingProvider {
	newProducer()
	return &KafkaMessagingProvider{}
}

func newProducer() {

	var err error
	var kafkaProducer *kafka.Producer

	if config.ENVRIONMENT_VARIABLES.KAFKA_PROTOCOL == "SASL_SSL" {
		kafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{
			"client.id":         name,
			"security.protocol": "sasl",
			"sasl.username":     config.ENVRIONMENT_VARIABLES.KAFKA_USERNAME,
			"sasl.password":     config.ENVRIONMENT_VARIABLES.KAFKA_PASSWORD,
			"sasl.mechanisms":   "plain",
			"bootstrap.servers": config.ENVRIONMENT_VARIABLES.KAFKA_BOOTSTRAP_SERVERS,
		})
	} else {
		kafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{
			"client.id":         name,
			"bootstrap.servers": config.ENVRIONMENT_VARIABLES.KAFKA_BOOTSTRAP_SERVERS,
		})
	}

	if err != nil {
		panic(err)
	}

	producer = kafkaProducer
}

func (KafkaMessagingProvider) Consume(topics []string, channel chan *kafka.Message) {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.ENVRIONMENT_VARIABLES.KAFKA_BOOTSTRAP_SERVERS,
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
