package main

import "github.com/gpbPiazza/alemao-bigodes/infra/http"

func main() {
	server := http.NewServer()

	defer server.Shutdown()

	server.Listen(":8080")
}
