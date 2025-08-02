#!/bin/bash

# Экспорт переменных окружения для базы данных
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=password
export DB_NAME=myproject

# Запуск приложения
echo "Starting Go application with database..."
go run main.go 