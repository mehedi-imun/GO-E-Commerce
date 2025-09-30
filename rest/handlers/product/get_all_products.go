package product

import (
	"ecommace/util"
	"net/http"
	"strconv"
	"sync"
)

var count int64

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

	products, err := h.service.GetAll(page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		dataCount, err := h.service.Count()
		if err != nil {
			http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
			return
		}
		count = dataCount
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		dataCount1, err := h.service.Count()
		if err != nil {
			http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
			return
		}
		count = dataCount1
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		dataCount2, err := h.service.Count()
		if err != nil {
			http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
			return
		}
		count = dataCount2
	}()
	wg.Wait()
	util.SendPage(w, products, limit, page, count)
}
