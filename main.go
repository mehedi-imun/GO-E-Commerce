package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

type Product struct {
	Id    int `json:"id"`
	Title string
}

var productLIst []Product

func productGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "get in invalid", 400)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(productLIst)

}

func main() {
	mux := http.NewServeMux() // router
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/product", productGet)
	fmt.Println("server is running on :3000") //route

	err := http.ListenAndServe(":3000", mux) // expose port

	if err != nil {
		fmt.Println("error", err) // error
	}

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
	productLIst = append(productLIst, pro1)
	productLIst = append(productLIst, pro2)
}
