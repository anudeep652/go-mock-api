package main

import (
	"github.com/anudeep652/go-mock-api/api"
)

func main() {
	app := api.Server()

	app.Listen(":3000")
}
