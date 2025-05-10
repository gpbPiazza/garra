package main

import "github.com/gpbPiazza/garra/infra/http"

func main() {
	server := http.NewServer()

	defer server.Shutdown()

	server.Listen(":8080")
}
