package database

import (
	"encoding/json"
	"net/http"
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
