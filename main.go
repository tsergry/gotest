package main

import (
	"fmt"
	"log"
	"net/http"

	"myproject/internal/database"
	"myproject/internal/handlers"
	"myproject/internal/models"
)

func main() {
	fmt.Println("Starting Go application...")

	// Подключение к базе данных
	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Инициализация репозиториев и обработчиков
	userRepo := models.NewUserRepository(db.DB)
	userHandler := handlers.NewUserHandler(userRepo)
	
	productRepo := models.NewProductRepository(db.DB)
	productHandler := handlers.NewProductHandler(productRepo)

	// Настройка маршрутов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go with PostgreSQL!")
	})

	// API маршруты для пользователей
	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetAllUsers(w, r)
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/users/get", userHandler.GetUser)
	http.HandleFunc("/api/users/update", userHandler.UpdateUser)
	http.HandleFunc("/api/users/delete", userHandler.DeleteUser)

	// API маршруты для товаров
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productHandler.GetAllProducts(w, r)
		case http.MethodPost:
			productHandler.CreateProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/products/get", productHandler.GetProduct)
	http.HandleFunc("/api/products/category", productHandler.GetProductsByCategory)
	http.HandleFunc("/api/products/update", productHandler.UpdateProduct)
	http.HandleFunc("/api/products/delete", productHandler.DeleteProduct)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
} 