package cmd

import (
	"ecommace/config"
	"ecommace/rest"
)

func Serve() {

	cnf := config.GetConfig()
	rest.Start(cnf)

}
