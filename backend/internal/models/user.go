package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *User) error {
	query := `
		INSERT INTO users (name, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	
	return r.db.QueryRow(query, user.Name, user.Email, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
}

func (r *UserRepository) GetByID(id int) (*User, error) {
	user := &User{}
	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1`
	
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	
	return user, nil
}

func (r *UserRepository) GetAll() ([]*User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users ORDER BY created_at DESC`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var users []*User
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	
	return users, nil
}

func (r *UserRepository) Update(user *User) error {
	query := `
		UPDATE users 
		SET name = $1, email = $2, updated_at = $3
		WHERE id = $4`
	
	user.UpdatedAt = time.Now()
	_, err := r.db.Exec(query, user.Name, user.Email, user.UpdatedAt, user.ID)
	return err
}

func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
} 