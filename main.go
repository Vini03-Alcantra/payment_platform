package main

import (
	"github.com/hyperyuri/webapi-with-go/datatase"
	"github.com/hyperyuri/webapi-with-go/server"
)

func main() {
	datatase.StartDB()
	server := server.NewServer()

	server.Run()
}
