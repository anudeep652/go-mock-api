package main

import (
	"fmt"

	"github.com/anudeep652/go-mock-api/api"
)

func main() {
	app := api.Server()
	go func() { fmt.Println(api.GenerateApiKey()) }()
	app.Listen(":3000")
}
