package models

import (
	"database/sql"
	"time"
)

type Order struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	ProductID   int       `json:"product_id"`
	Quantity    int       `json:"quantity"`
	TotalPrice  float64   `json:"total_price"`
	Status      string    `json:"status"` // pending, confirmed, shipped, delivered, cancelled
	OrderDate   time.Time `json:"order_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type OrderWithDetails struct {
	Order
	UserName    string  `json:"user_name"`
	UserEmail   string  `json:"user_email"`
	ProductName string  `json:"product_name"`
	UnitPrice   float64 `json:"unit_price"`
}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *Order) error {
	query := `
		INSERT INTO orders (user_id, product_id, quantity, total_price, status, order_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`
	
	now := time.Now()
	order.OrderDate = now
	order.CreatedAt = now
	order.UpdatedAt = now
	
	return r.db.QueryRow(query, order.UserID, order.ProductID, order.Quantity, 
		order.TotalPrice, order.Status, order.OrderDate, order.CreatedAt, order.UpdatedAt).Scan(&order.ID)
}

func (r *OrderRepository) GetByID(id int) (*OrderWithDetails, error) {
	order := &OrderWithDetails{}
	query := `
		SELECT o.id, o.user_id, o.product_id, o.quantity, o.total_price, o.status, 
		       o.order_date, o.created_at, o.updated_at,
		       u.name as user_name, u.email as user_email,
		       p.name as product_name, p.price as unit_price
		FROM orders o
		JOIN users u ON o.user_id = u.id
		JOIN products p ON o.product_id = p.id
		WHERE o.id = $1`
	
	err := r.db.QueryRow(query, id).Scan(&order.ID, &order.UserID, &order.ProductID, 
		&order.Quantity, &order.TotalPrice, &order.Status, &order.OrderDate, 
		&order.CreatedAt, &order.UpdatedAt, &order.UserName, &order.UserEmail, 
		&order.ProductName, &order.UnitPrice)
	if err != nil {
		return nil, err
	}
	
	return order, nil
}

func (r *OrderRepository) GetAll() ([]*OrderWithDetails, error) {
	query := `
		SELECT o.id, o.user_id, o.product_id, o.quantity, o.total_price, o.status, 
		       o.order_date, o.created_at, o.updated_at,
		       u.name as user_name, u.email as user_email,
		       p.name as product_name, p.price as unit_price
		FROM orders o
		JOIN users u ON o.user_id = u.id
		JOIN products p ON o.product_id = p.id
		ORDER BY o.created_at DESC`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var orders []*OrderWithDetails
	for rows.Next() {
		order := &OrderWithDetails{}
		err := rows.Scan(&order.ID, &order.UserID, &order.ProductID, 
			&order.Quantity, &order.TotalPrice, &order.Status, &order.OrderDate, 
			&order.CreatedAt, &order.UpdatedAt, &order.UserName, &order.UserEmail, 
			&order.ProductName, &order.UnitPrice)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	
	return orders, nil
}

func (r *OrderRepository) GetByUserID(userID int) ([]*OrderWithDetails, error) {
	query := `
		SELECT o.id, o.user_id, o.product_id, o.quantity, o.total_price, o.status, 
		       o.order_date, o.created_at, o.updated_at,
		       u.name as user_name, u.email as user_email,
		       p.name as product_name, p.price as unit_price
		FROM orders o
		JOIN users u ON o.user_id = u.id
		JOIN products p ON o.product_id = p.id
		WHERE o.user_id = $1
		ORDER BY o.created_at DESC`
	
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var orders []*OrderWithDetails
	for rows.Next() {
		order := &OrderWithDetails{}
		err := rows.Scan(&order.ID, &order.UserID, &order.ProductID, 
			&order.Quantity, &order.TotalPrice, &order.Status, &order.OrderDate, 
			&order.CreatedAt, &order.UpdatedAt, &order.UserName, &order.UserEmail, 
			&order.ProductName, &order.UnitPrice)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	
	return orders, nil
}

func (r *OrderRepository) UpdateStatus(id int, status string) error {
	query := `UPDATE orders SET status = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.Exec(query, status, time.Now(), id)
	return err
}

func (r *OrderRepository) Delete(id int) error {
	query := `DELETE FROM orders WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *OrderRepository) GetByStatus(status string) ([]*OrderWithDetails, error) {
	query := `
		SELECT o.id, o.user_id, o.product_id, o.quantity, o.total_price, o.status, 
		       o.order_date, o.created_at, o.updated_at,
		       u.name as user_name, u.email as user_email,
		       p.name as product_name, p.price as unit_price
		FROM orders o
		JOIN users u ON o.user_id = u.id
		JOIN products p ON o.product_id = p.id
		WHERE o.status = $1
		ORDER BY o.created_at DESC`
	
	rows, err := r.db.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var orders []*OrderWithDetails
	for rows.Next() {
		order := &OrderWithDetails{}
		err := rows.Scan(&order.ID, &order.UserID, &order.ProductID, 
			&order.Quantity, &order.TotalPrice, &order.Status, &order.OrderDate, 
			&order.CreatedAt, &order.UpdatedAt, &order.UserName, &order.UserEmail, 
			&order.ProductName, &order.UnitPrice)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	
	return orders, nil
} 