package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

// Шаблон для отображения HTML-страницы
var tmpl = template.Must(template.New("form").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Форма ввода</title>
</head>
<body>
    <h1>Введите данные</h1>
    <form method="POST" action="/">
        <label for="name">Имя:</label>
        <input type="text" id="name" name="name" required><br><br>

        <label for="age">Возраст:</label>
        <input type="number" id="age" name="age" required><br><br>

        <label for="email">Email:</label>
        <input type="email" id="email" name="email" required><br><br>

        <label for="carNumber">Номер автомобиля:</label>
        <input type="text" id="carNumber" name="carNumber" placeholder="Например, A123BC" required><br><br>

        <button type="submit">Отправить</button>
    </form>
</body>
</html>
`))

// Шаблон для отображения результата
var resultTmpl = template.Must(template.New("result").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Результат</title>
</head>
<body>
    <h1>Введенные данные</h1>
    <p><strong>Имя:</strong> {{.Name}}</p>
    <p><strong>Возраст:</strong> {{.Age}}</p>
    <p><strong>Email:</strong> {{.Email}}</p>
    <p><strong>Номер автомобиля:</strong> {{.CarNumber}}</p>
</body>
</html>
`))

// Структура для хранения данных
type FormData struct {
	Name      string
	Age       int
	Email     string
	CarNumber string
}

// Главная страница (GET)
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Отображаем форму
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		// Обрабатываем данные формы
		r.ParseForm()

		// Получаем данные из формы
		name := r.FormValue("name")
		ageStr := r.FormValue("age")
		email := r.FormValue("email")
		carNumber := r.FormValue("carNumber")

		// Проверяем корректность данных
		age, err := strconv.Atoi(ageStr)
		if err != nil || age <= 0 {
			http.Error(w, "Возраст должен быть положительным числом", http.StatusBadRequest)
			return
		}

		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(email) {
			http.Error(w, "Некорректный email", http.StatusBadRequest)
			return
		}

		carNumberRegex := regexp.MustCompile(`^[A-ZА-Я]{1}\d{3}[A-ZА-Я]{2}$`)
		if !carNumberRegex.MatchString(carNumber) {
			http.Error(w, "Некорректный номер автомобиля (пример: A123BC)", http.StatusBadRequest)
			return
		}

		// Формируем данные для отображения
		data := FormData{
			Name:      name,
			Age:       age,
			Email:     email,
			CarNumber: carNumber,
		}

		// Отображаем результат
		resultTmpl.Execute(w, data)
	}
}

func main() {
	// Регистрируем обработчик
	http.HandleFunc("/", formHandler)

	// Запускаем сервер
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
