package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/bubbletea"
	"github.com/segmentio/kafka-go"
)

// KafkaConfig содержит конфигурацию для Kafka
type KafkaConfig struct {
	Broker  string
	Topic   string
	GroupID string
}

// Model для bubbletea
type Model struct {
	messages []string
	writer   *kafka.Writer
	reader   *kafka.Reader
}

// Init инициализирует модель
func (m Model) Init() tea.Cmd {
	return nil
}

// Update обрабатывает сообщения и обновляет модель
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case string:
		m.messages = append(m.messages, msg)
		return m, nil
	}
	return m, nil
}

// View отображает текущее состояние модели
func (m Model) View() string {
	s := "Сообщения:\n\n"
	for _, msg := range m.messages {
		s += fmt.Sprintf("  %s\n", msg)
	}
	s += "\nНажмите q для выхода.\n"
	return s
}

// StartKafkaProducer запускает продюсера Kafka
func StartKafkaProducer(writer *kafka.Writer, msgChan chan string) {
	for {
		select {
		case msg := <-msgChan:
			err := writer.WriteMessages(context.Background(), kafka.Message{
				Value: []byte(msg),
			})
			if err != nil {
				log.Printf("Ошибка отправки сообщения: %v\n", err)
			} else {
				log.Printf("Сообщение отправлено: %s\n", msg)
			}
		}
	}
}

// StartKafkaConsumer запускает консьюмера Kafka
func StartKafkaConsumer(reader *kafka.Reader, msgChan chan string) {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Ошибка получения сообщения: %v\n", err)
			continue
		}
		log.Printf("Сообщение получено: %s\n", string(msg.Value))
		msgChan <- string(msg.Value)
	}
}

func main() {
	// Конфигурация Kafka
	config := KafkaConfig{
		Broker:  "localhost:9092",
		Topic:   "test-topic",
		GroupID: "test-group",
	}

	// Создание продюсера Kafka
	writer := &kafka.Writer{
		Addr:     kafka.TCP(config.Broker),
		Topic:    config.Topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	// Создание консьюмера Kafka
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{config.Broker},
		Topic:    config.Topic,
		GroupID:  config.GroupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e4, // 100KB
	})
	defer reader.Close()

	// Каналы для сообщений
	msgChan := make(chan string)
	kafkaMsgChan := make(chan string)

	// Запуск продюсера и консьюмера
	go StartKafkaProducer(writer, msgChan)
	go StartKafkaConsumer(reader, kafkaMsgChan)

	// Инициализация модели bubbletea
	model := Model{
		writer: writer,
		reader: reader,
	}

	// Запуск bubbletea
	p := tea.NewProgram(model)

	// Горутина для обработки сообщений из Kafka
	go func() {
		for {
			select {
			case msg := <-kafkaMsgChan:
				p.Send(msg)
			}
		}
	}()

	// Горутина для отправки сообщений в Kafka
	go func() {
		for {
			var input string
			fmt.Print("Введите сообщение: ")
			fmt.Scanln(&input)
			msgChan <- input
		}
	}()

	// Ожидание сигнала завершения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Завершение программы
	fmt.Println("Завершение программы...")
}
