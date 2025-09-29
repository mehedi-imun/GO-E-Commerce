package product

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// GetProductByID handles GET /products/{id}
func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	// extract id from URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Product ID must be an integer", http.StatusBadRequest)
		return
	}

	product, err := h.service.FindByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
