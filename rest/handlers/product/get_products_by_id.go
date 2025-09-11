package product

import (
	"ecommace/database"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h*Handler)GetProductByID(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)
	if err != nil {
		return
	}
	for _, product := range database.ProductLIst {
		if product.Id == pId {
			encoder := json.NewEncoder(w)
			encoder.Encode(product)
			return
		}

	}
	encoder := json.NewEncoder(w)
	encoder.Encode("data not found")

}