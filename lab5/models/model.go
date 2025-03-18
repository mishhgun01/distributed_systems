package models

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

// KafkaConfig содержит конфигурацию для Kafka
type KafkaConfig struct {
	Broker  string
	Topic   string
	GroupID string
}

// StartKafkaProducer запускает продюсера Kafka
func StartKafkaProducer(writer *kafka.Writer) {
	for {
		var input string
		fmt.Print("Введите сообщение: ")
		fmt.Scanln(&input)
		err := writer.WriteMessages(context.Background(), kafka.Message{
			Value: []byte(input),
		})
		if err != nil {
			log.Printf("Ошибка отправки сообщения: %v\n", err)
		} else {
			log.Printf("Сообщение отправлено: %s\n", input)
		}

	}
}

// StartKafkaConsumer запускает консьюмера Kafka
func StartKafkaConsumer(reader *kafka.Reader) {
	log.Println("consumer started")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Ошибка получения сообщения: %v\n", err)
			continue
		}
		log.Printf("Сообщение получено: %s\n", string(msg.Value))
	}
}
