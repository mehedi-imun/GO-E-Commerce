package product

import (
	"ecommace/util"
	"net/http"
	"strconv"
)

// GetAllProducts handles GET /products
func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()
	pageStr := reqQuery.Get("page")
	limitStr := reqQuery.Get("limit")
	page, _ := strconv.ParseInt(pageStr, 10, 32)
	limit, _ := strconv.ParseInt(limitStr, 10, 32)
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	products, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	util.SendPage(w, products, limit, page, 0)
}
