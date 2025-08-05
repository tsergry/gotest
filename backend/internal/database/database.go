package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewConnection() (*DB, error) {
	// Получаем параметры подключения из переменных окружения
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "orders_db")

	// Извлекаем хост и порт из DB_HOST
	var dbHost, dbPort string
	if strings.Contains(host, ":") {
		parts := strings.Split(host, ":")
		dbHost = parts[0]
		dbPort = parts[1]
	} else {
		dbHost = host
		dbPort = port
	}

	// Формируем строку подключения
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, user, password, dbname)

	// Открываем соединение с базой данных
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Проверяем соединение
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	log.Println("Successfully connected to database")
	return &DB{db}, nil
}

func (db *DB) Close() error {
	return db.DB.Close()
}

// getEnv получает значение переменной окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
