package product

import (
	"ecommace/database"
	"encoding/json"
	"net/http"
)

// GetProducts returns all products
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database.ProductLIst)
}
