package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	Description string  `db:"description" json:"description"`
	Price       float64 `db:"price" json:"price"`
	Stock       int     `db:"stock" json:"stock"`
	Image_Url   string  `db:"description" json:"Image_Url"`
}

type ProductRepo interface {
	Create(product Product) (*Product, error)
	FindByID(id int) (*Product, error)
	GetAll() ([]*Product, error)
	UpdateStock(id, stock int) error
	Delete(id int) error
}

type productRepo struct {
	db *sqlx.DB
}

// constructor
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{db: db}
}

// Create product
func (r *productRepo) Create(product Product) (*Product, error) {
	query := `
		INSERT INTO products (name, description, price, stock)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	err := r.db.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
	).Scan(&product.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to insert product: %w", err)
	}

	return &product, nil
}

// Find product by ID
func (r *productRepo) FindByID(id int) (*Product, error) {
	var p Product
	query := `SELECT id, name, description, price, stock FROM products WHERE id=$1`

	err := r.db.Get(&p, query, id)
	if err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	return &p, nil
}

// Get all products
func (r *productRepo) GetAll() ([]*Product, error) {
	var products []*Product
	query := `SELECT id, name, description, price, stock FROM products`

	err := r.db.Select(&products, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	return products, nil
}

// Update stock
func (r *productRepo) UpdateStock(id, stock int) error {
	query := `UPDATE products SET stock=$1 WHERE id=$2`
	_, err := r.db.Exec(query, stock, id)
	if err != nil {
		return fmt.Errorf("failed to update stock: %w", err)
	}
	return nil
}

// Delete product
func (r *productRepo) Delete(id int) error {
	query := `DELETE FROM products WHERE id=$1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}
