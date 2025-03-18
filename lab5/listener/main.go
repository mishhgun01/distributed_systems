package main

import (
	"fmt"
	"lab5/models"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Конфигурация Kafka
	config := models.KafkaConfig{
		Broker:  "localhost:9092",
		Topic:   "test-topic-1",
		GroupID: "test-group",
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{config.Broker},
		Topic:    config.Topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e4, // 100KB
	})
	defer reader.Close()

	go models.StartKafkaConsumer(reader)

	// Ожидание сигнала завершения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Завершение программы
	fmt.Println("Завершение слушателя...")
}
