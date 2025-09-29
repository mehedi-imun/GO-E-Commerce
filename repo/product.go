package repo

import (
	"ecommace/domain"
	"ecommace/product"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

// constructor
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{db: db}
}

// Create product
func (r *productRepo) Create(product domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO products (name, description, price, stock, Image_Url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := r.db.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.Image_Url,
	).Scan(&product.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to insert product: %w", err)
	}

	return &product, nil
}

// Find product by ID
func (r *productRepo) FindByID(id int) (*domain.Product, error) {
	var p domain.Product
	query := `SELECT id, name, description, price, stock FROM products WHERE id=$1`

	err := r.db.Get(&p, query, id)
	if err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	return &p, nil
}

// Get all products
func (r *productRepo) GetAll() ([]*domain.Product, error) {
	var products []*domain.Product
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
