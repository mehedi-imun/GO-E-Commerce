package database

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Product struct {
	Id    int `json:"id"`
	Title string
}

var ProductLIst []Product

func ProductGet(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(ProductLIst)

}
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)
	if err != nil {
		return
	}
	for _, product := range ProductLIst {
		if product.Id == pId {
			encoder := json.NewEncoder(w)
			encoder.Encode(product)
			return
		}

	}
	encoder := json.NewEncoder(w)
	encoder.Encode("data not found")

}
func init() {
	pro1 := Product{
		Id:    1,
		Title: "Hello",
	}
	pro2 := Product{
		Id:    2,
		Title: "Hello",
	}
	ProductLIst = append(ProductLIst, pro1)
	ProductLIst = append(ProductLIst, pro2)
}
