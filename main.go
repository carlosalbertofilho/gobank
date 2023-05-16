package main

import (
	"github.com/carlosalbertofilho/gobank/api"
)

func main() {
	server := api.NewAPIServer(":8080")
	server.Run()
}
