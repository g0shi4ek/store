# Этап сборки
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Установка migrate и сборка приложения
RUN wget -O migrate.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz && \
    tar -xzf migrate.tar.gz && \
    mv migrate /usr/local/bin/

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/main ./cmd/main.go

# Этап запуска
FROM alpine:latest
WORKDIR /app

# Копируем только необходимое
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/

CMD ["sh", "-c", "\
    if [ \"$RUN_MIGRATIONS\" = \"true\" ]; then \
        migrate -path ./migrations -database \"$DATABASE_URL\" up; \
    fi; \
    ./main \
"]

EXPOSE 8080