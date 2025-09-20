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
	GetAll() ([]*User, error)
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
		INSERT INTO users (
		first_name, 
		last_name, 
		email, 
		password, 
		is_shop_owner
		)VALUES (
		$1, 
		$2, 
		$3, 
		$4, 
		$5
		)
		RETURNING id
	`

	// Use NamedQueryRow for struct binding and returning id
	row := r.db.QueryRow(query, user.FirstName, user.FirstName, user.Email, user.Password, user.IsShopOwner)
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}

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

func (r *userRepo) GetAll() ([]*User, error) {
	var users []*User
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner FROM users`
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
