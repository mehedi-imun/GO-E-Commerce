package product

import (
	"encoding/json"
	"net/http"
)

// GetAllProducts handles GET /products
func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
