package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ecommace/repo"
)

// CreateProduct handles POST /products
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p repo.Product

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	created, err := h.repo.Create(p)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create product: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(created)
}
