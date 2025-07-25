package minuta

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate_many_buyers_to_one_seller(t *testing.T) {
	t.Parallel()
	generatorApp := NewGeneratorApp()

	expected := `
	<fragmento indice="CABECALHO" />
	<br>
	<u>TRANSMITENTE(S)</u>:> <span><strong>ADMINISTRADORA DE BENS CATARINA LTDA</strong>., supraqualificada.</span>
	<br/>
	<u>ADQUIRENTE(S)</u>:><span><strong>MÁRCIO JEAN RICARDO</strong>, brasileiro, policial militar, CPF nº 939.457.799-87, e sua esposa <strong>ELISÂNGELA SILVA LOPES RICARDO</strong>, brasileira, técnica de laboratório, CPF n° 049.369.096-45, casados pelo regime da Comunhão Parcial de Bens, na vigência da Lei n° 6.515/77, residentes e domiciliados na Rua Adolfo Gleich, n° 21, Bairro Jardim Maluche, Brusque/SC.</span>
	<br>
	<u>FORMA DO TÍTULO</u>: Escritura Pública de Compra e Venda, lavrada pelo 1º Tabelionato de Notas e Protesto de Brusque/SC, Livro 968, Folhas 024/026V, em 28/04/2025.
	<br/><u>VALOR</u>: R$ 303.000,00 (trezentos e três mil reais).
	<br/><u>CONDIÇÕES</u>: Não constam.
	<br/><u>OBSERVAÇÕES</u>:
	<strong>ITBI</strong>: Recolhido no valor de R$ 6.134,16, com incidência sobre R$ 303.000,00, devidamente quitado. No ato da lavratura da Escritura Pública, foram apresentadas as certidões previstas em Lei. Com as demais cláusulas e condições da Escritura Pública.<strong> NO PRAZO REGULAMENTAR SERÁ EMITIDA A DOI</strong>.
	<fragmento indice="FINALIZACAO_ATO" />.
	`

	doc, err := os.ReadFile("../../infra/test_files/ato_consultar_tjsc_many_buyer_to_one_seller.txt")
	require.NoError(t, err)

	params := GenerateParams{
		DocStr:               string(doc),
		IsTransOverqualified: true,
		IsAdquiOverqualified: false,
	}

	got, err := generatorApp.Generate(params)
	require.NoError(t, err)

	assert.Equal(t, expected, got.MinutaHTML)
}
