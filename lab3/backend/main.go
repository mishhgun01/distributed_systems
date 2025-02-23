package main

import (
	"log"
	"net/http"
	"os"

	"notes/handlers"
	"notes/models"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Подключение к базе данных
	dsn := "host=db user=postgres password=postgres dbname=notes port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	// Миграция модели
	db.AutoMigrate(&models.Note{})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Инициализация маршрутов
	r := mux.NewRouter()
	r.HandleFunc("/api/notes", handlers.GetNotes(db)).Methods(http.MethodGet)
	r.HandleFunc("/api/notes", handlers.UpdateNote(db)).Methods(http.MethodPut)
	r.HandleFunc("/api/notes", handlers.CreateNote(db)).Methods(http.MethodPost)
	r.HandleFunc("/api/notes/{id:[0-9]+}", handlers.DeleteNote(db)).Methods(http.MethodDelete)
	handler := c.Handler(r)

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Сервер запущен на порту %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
