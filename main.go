package main

import (
	"github.com/briheet/goKafka/api"
)

func main() {
	println("Hello from goKafka")

	api := api.NewApiServer(":3000")
	api.Serve()
}
