package main

import (
	"log"
	"os"

	"github.com/briheet/goKafka/api"
)

func main() {
	println("Hello from goKafka")

	api, err := api.NewApiServer("3000")
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	api.Serve()
}
