# Этап сборки
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Установка Go-Migrate
RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz && \
    tar -xvzf migrate.linux-amd64.tar.gz && \
    mv migrate /usr/local/bin/

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/main ./cmd/main.go

# Этап запуска
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate
COPY ./migrations ./migrations
COPY .env ./
EXPOSE 8080
CMD ["/app/main"]