package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"myproject/internal/models"
)

type OrderHandler struct {
	orderRepo   *models.OrderRepository
	productRepo *models.ProductRepository
	userRepo    *models.UserRepository
}

func NewOrderHandler(orderRepo *models.OrderRepository, productRepo *models.ProductRepository, userRepo *models.UserRepository) *OrderHandler {
	return &OrderHandler{
		orderRepo:   orderRepo,
		productRepo: productRepo,
		userRepo:    userRepo,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Проверяем существование пользователя
	_, err := h.userRepo.GetByID(order.UserID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	// Проверяем существование товара
	product, err := h.productRepo.GetByID(order.ProductID)
	if err != nil {
		http.Error(w, "Product not found", http.StatusBadRequest)
		return
	}

	// Проверяем наличие товара на складе
	if product.Stock < order.Quantity {
		http.Error(w, "Insufficient stock", http.StatusBadRequest)
		return
	}

	// Рассчитываем общую стоимость
	order.TotalPrice = product.Price * float64(order.Quantity)
	order.Status = "pending"

	if err := h.orderRepo.Create(&order); err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := h.orderRepo.GetByID(id)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	orders, err := h.orderRepo.GetAll()
	if err != nil {
		http.Error(w, "Failed to get orders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) GetOrdersByUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	orders, err := h.orderRepo.GetByUserID(userID)
	if err != nil {
		http.Error(w, "Failed to get user orders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) GetOrdersByStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	status := r.URL.Query().Get("status")
	if status == "" {
		http.Error(w, "Status parameter is required", http.StatusBadRequest)
		return
	}

	orders, err := h.orderRepo.GetByStatus(status)
	if err != nil {
		http.Error(w, "Failed to get orders by status", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	var statusUpdate struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&statusUpdate); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Проверяем валидность статуса
	validStatuses := []string{"pending", "confirmed", "shipped", "delivered", "cancelled"}
	isValid := false
	for _, status := range validStatuses {
		if status == statusUpdate.Status {
			isValid = true
			break
		}
	}
	if !isValid {
		http.Error(w, "Invalid status", http.StatusBadRequest)
		return
	}

	if err := h.orderRepo.UpdateStatus(id, statusUpdate.Status); err != nil {
		http.Error(w, "Failed to update order status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order status updated successfully"})
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	if err := h.orderRepo.Delete(id); err != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
} 