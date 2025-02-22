# Этап сборки
FROM golang:1.23 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY go.mod  ./
RUN go mod download

COPY . .

# Сборка приложения для Linux
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Этап запуска
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Устанавливаем необходимые библиотеки
RUN apk add --no-cache libc6-compat

# Копируем скомпилированное приложение
COPY --from=builder /app/main .

# Делаем файл исполняемым
RUN chmod +x ./main

# Указываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]