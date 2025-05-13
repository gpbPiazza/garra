package minuta

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCabecalho(t *testing.T) {
	t.Run("Compra e venda type R good params", func(t *testing.T) {
		params := CabecalhoParams{
			AtoType:       "R",
			TitleAto:      "COMPRE E VENDA",
			NumAto:        "10",
			Matricula:     "50.679",
			RegistroDate:  time.Date(2025, 04, 07, 0, 0, 0, 0, time.UTC),
			Protocolo:     "253.426",
			ProtocoloDate: time.Date(2025, 04, 01, 0, 0, 0, 0, time.UTC),
		}

		expected := `
<cabecalho>
  <b>R.10-50.679, </b> em 7 de April de 2025.
  <br>
  <b>Prot. 253.426,</b> datado de 01/04/2025.
  <br>
  <u><b><maiusculo>COMPRE E VENDA</maiusculo></b></u>:
</cabecalho>
`
		got := Cabecalho(params)

		assert.Equal(t, expected, got)
	})
}
