package main

import (
	"fmt"

	"github.com/gpbPiazza/garra/infra/envs"
	"github.com/gpbPiazza/garra/infra/http"
)

func main() {
	env := envs.GetEnvs()

	server := http.NewServer()

	defer server.Shutdown()

	server.Listen(fmt.Sprintf(":%s", env.ApiPort))
}
