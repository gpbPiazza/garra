package main

import (
	"fmt"
	"time"

	"github.com/gpbPiazza/alemao-bigodes/minuta"
)

func main() {
	params := minuta.CabecalhoParams{
		AtoType:       "R",
		TitleAto:      "COMPRE E VENDA",
		NumAto:        "10",
		Matricula:     "50.679",
		RegistroDate:  time.Date(2025, 04, 07, 0, 0, 0, 0, time.UTC),
		Protocolo:     "253.426",
		ProtocoloDate: time.Date(2025, 04, 01, 0, 0, 0, 0, time.UTC),
	}

	got := minuta.Cabecalho(params)

	fmt.Println(got)
}
