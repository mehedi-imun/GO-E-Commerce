package product

import (
	"ecommace/database"
	"encoding/json"
	"net/http"
)

// CreateProduct adds a new product
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p database.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	p.Id = len(database.ProductLIst) + 1
	database.ProductLIst = append(database.ProductLIst, p)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
