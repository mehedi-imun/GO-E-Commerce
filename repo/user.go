package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `db:"id" json:"id"`
	FirstName   string `db:"first_name" json:"first_name"`
	LastName    string `db:"last_name" json:"last_name"`
	Email       string `db:"email" json:"email"`
	Password    string `db:"password" json:"password"`
	IsShopOwner bool   `db:"is_shop_owner" json:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

// Constructor
func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{db: db}
}

// Create inserts a new user and returns it
func (r *userRepo) Create(user User) (*User, error) {
	query := `
		INSERT INTO users (first_name, last_name, email, password, is_shop_owner)
		VALUES (:first_name, :last_name, :email, :password, :is_shop_owner)
		RETURNING id
	`

	// Use NamedQueryRow for struct binding and returning id
	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	var id int
	if err := stmt.Get(&id, user); err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	user.ID = id
	return &user, nil
}

// Find retrieves a user by email and password
func (r *userRepo) Find(email, pass string) (*User, error) {
	var u User
	query := `
		SELECT id, first_name, last_name, email, password, is_shop_owner
		FROM users
		WHERE email = $1 AND password = $2
	`

	err := r.db.Get(&u, query, email, pass)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return &u, nil
}
