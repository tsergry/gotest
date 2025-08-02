package models

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *Product) error {
	query := `
		INSERT INTO products (name, description, price, category, stock, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`
	
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now
	
	return r.db.QueryRow(query, product.Name, product.Description, product.Price, 
		product.Category, product.Stock, product.CreatedAt, product.UpdatedAt).Scan(&product.ID)
}

func (r *ProductRepository) GetByID(id int) (*Product, error) {
	product := &Product{}
	query := `SELECT id, name, description, price, category, stock, created_at, updated_at 
			  FROM products WHERE id = $1`
	
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, 
		&product.Price, &product.Category, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}
	
	return product, nil
}

func (r *ProductRepository) GetAll() ([]*Product, error) {
	query := `SELECT id, name, description, price, category, stock, created_at, updated_at 
			  FROM products ORDER BY created_at DESC`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var products []*Product
	for rows.Next() {
		product := &Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, 
			&product.Price, &product.Category, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	
	return products, nil
}

func (r *ProductRepository) Update(product *Product) error {
	query := `
		UPDATE products 
		SET name = $1, description = $2, price = $3, category = $4, stock = $5, updated_at = $6
		WHERE id = $7`
	
	product.UpdatedAt = time.Now()
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, 
		product.Category, product.Stock, product.UpdatedAt, product.ID)
	return err
}

func (r *ProductRepository) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ProductRepository) GetByCategory(category string) ([]*Product, error) {
	query := `SELECT id, name, description, price, category, stock, created_at, updated_at 
			  FROM products WHERE category = $1 ORDER BY created_at DESC`
	
	rows, err := r.db.Query(query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var products []*Product
	for rows.Next() {
		product := &Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, 
			&product.Price, &product.Category, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	
	return products, nil
} 