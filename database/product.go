package database

type Product struct {
	Id    int `json:"id"`
	Title string
}

var ProductLIst []Product

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
