package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
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
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	// Проверяем соединение
	if err = db.Ping(); err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	log.Println("Successfully connected to database")

	// SQL скрипт для создания таблиц
	sqlScript := `
	-- Создание таблицы пользователей
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	-- Создание индексов для оптимизации запросов
	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	CREATE INDEX IF NOT EXISTS idx_users_created_at ON users(created_at);

	-- Вставка тестовых данных
	INSERT INTO users (name, email) VALUES 
		('John Doe', 'john@example.com'),
		('Jane Smith', 'jane@example.com'),
		('Bob Johnson', 'bob@example.com')
	ON CONFLICT (email) DO NOTHING;

	-- Создание таблицы товаров
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		price DECIMAL(10,2) NOT NULL,
		category VARCHAR(100) NOT NULL,
		stock INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	-- Создание индексов для товаров
	CREATE INDEX IF NOT EXISTS idx_products_category ON products(category);
	CREATE INDEX IF NOT EXISTS idx_products_created_at ON products(created_at);
	CREATE INDEX IF NOT EXISTS idx_products_price ON products(price);

	-- Вставка тестовых товаров
	INSERT INTO products (name, description, price, category, stock) VALUES 
		('iPhone 15', 'Смартфон Apple iPhone 15 128GB', 89999.99, 'Electronics', 10),
		('MacBook Pro', 'Ноутбук Apple MacBook Pro 14" M3', 249999.99, 'Electronics', 5),
		('Nike Air Max', 'Кроссовки Nike Air Max 270', 12999.99, 'Shoes', 20),
		('Coffee Maker', 'Кофемашина Philips 2200', 15999.99, 'Home', 8),
		('Book: Go Programming', 'Учебник по программированию на Go', 2999.99, 'Books', 15)
	ON CONFLICT DO NOTHING;

	-- Создание таблицы заказов
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES users(id),
		product_id INTEGER NOT NULL REFERENCES products(id),
		quantity INTEGER NOT NULL,
		total_price DECIMAL(10,2) NOT NULL,
		status VARCHAR(50) NOT NULL DEFAULT 'pending',
		order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	-- Создание индексов для заказов
	CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id);
	CREATE INDEX IF NOT EXISTS idx_orders_product_id ON orders(product_id);
	CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);
	CREATE INDEX IF NOT EXISTS idx_orders_created_at ON orders(created_at);

	-- Вставка тестовых заказов
	INSERT INTO orders (user_id, product_id, quantity, total_price, status) VALUES 
		(1, 1, 1, 89999.99, 'pending'),
		(2, 3, 2, 25999.98, 'confirmed'),
		(1, 4, 1, 15999.99, 'shipped'),
		(3, 2, 1, 249999.99, 'delivered'),
		(2, 5, 3, 8999.97, 'cancelled')
	ON CONFLICT DO NOTHING;
	`

	// Выполняем SQL скрипт
	_, err = db.Exec(sqlScript)
	if err != nil {
		log.Fatal("Error executing SQL script:", err)
	}

	log.Println("Database initialized successfully!")
}

// getEnv получает значение переменной окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
