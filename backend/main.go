package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/models"
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

	orderRepo := models.NewOrderRepository(db.DB)
	orderHandler := handlers.NewOrderHandler(orderRepo, productRepo, userRepo)

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

	// API маршруты для заказов
	http.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			orderHandler.GetAllOrders(w, r)
		case http.MethodPost:
			orderHandler.CreateOrder(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/orders/get", orderHandler.GetOrder)
	http.HandleFunc("/api/orders/user", orderHandler.GetOrdersByUser)
	http.HandleFunc("/api/orders/status", orderHandler.GetOrdersByStatus)
	http.HandleFunc("/api/orders/update-status", orderHandler.UpdateOrderStatus)
	http.HandleFunc("/api/orders/delete", orderHandler.DeleteOrder)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
