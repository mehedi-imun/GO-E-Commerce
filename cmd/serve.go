package cmd

import (
	"ecommace/config"
	"ecommace/rest"
	"ecommace/rest/handlers/product"
)

func Serve() {

	cnf := config.GetConfig()
	productHandler:= product.NewHandler()
	server:= rest.NewServer(productHandler)
	server.Start(cnf)

}
