package main

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"lab5/models"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Конфигурация Kafka
	config := models.KafkaConfig{
		Broker: "localhost:9092",
		Topic:  "test-topic-1",
	}

	// Создание продюсера Kafka
	writer := &kafka.Writer{
		Addr:        kafka.TCP(config.Broker),
		Topic:       config.Topic,
		Balancer:    &kafka.LeastBytes{},
		MaxAttempts: 10,
	}
	defer writer.Close()
	// Запуск продюсера и консьюмера
	go models.StartKafkaProducer(writer)
	// Горутина для отправки сообщений в Kafka

	// Ожидание сигнала завершения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Завершение программы
	fmt.Println("Завершение писателя...")
}
