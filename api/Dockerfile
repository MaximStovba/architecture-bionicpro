FROM golang:1.24-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Установка зависимостей для сборки Go-проектов
RUN apk add --no-cache git build-base

# Кэшируем зависимости Go
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка
RUN go build -o main .

# Открываем порт
EXPOSE 8000

# Запуск
CMD ["./main"]
