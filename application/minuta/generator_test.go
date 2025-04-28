package minuta

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate_one_to_one_buy_CNPJ_and_sell_CPF(t *testing.T) {
	generatorApp := NewGeneratorApp()

	expected := `
<fragmento indice="CABECALHO" />
<br>
<u>TRANSMITENTE(S)</u>:> <span>CONSTRUARTE CONSTRUTORA E INCORPORADORA LTDA, brasileiro, Casado (a), CPF nº 29.235.977/0001-70, residente e domiciliado na Rua Luiz Eccel, nº 1, Bairro Paquetá, Brusque/SC.</span>.
<br/>
<u>ADQUIRENTE(S)</u>:><span>SIDNEI ANTONIO GATTIS, brasileiro, Solteiro (a), CPF nº 037.561.669-10, residente e domiciliado na Rua Azambuja, nº 541, Bairro Azambuja, Brusque/SC.</span>.
<br>
<u>FORMA DO TÍTULO</u>: Escritura Pública de Compra e Venda, lavrada pelo 1º Tabelionato de Notas e de Protesto de Brusque/SC, Livro 965, Folhas 121/123V, em 26/03/2025. 
<br/><u>VALOR</u>: R$ 120.000,00 (cento e vinte mil reais).
<br/><u>CONDIÇÕES</u>: Não constam.
<br/><u>OBSERVAÇÕES</u>:
<strong>ITBI</strong>: Recolhido no valor de R$ 2.900,00, com incidência sobre R$ 145.000,00, devidamente quitado. Nos termos do artigo 320 do CNCGFE/SC, o imóvel da presente matrícula, teve como valor atribuído de mercado, no
''quantum'' de R$ XXXXX. No ato da lavratura da Escritura Pública, foram apresentadas as certidões previstas em Lei. Com as demais cláusulas e condições da Escritura Pública. <strong> NO PRAZO REGULAMENTAR SERÁ EMITIDA A DOI</strong>.
<fragmento indice="FINALIZACAO_ATO" />.
`

	doc, err := os.ReadFile("../../infra/test_files/ato_consultar_tjsc_1_to_1_buy_and_sell.txt")
	require.NoError(t, err)

	got, err := generatorApp.Generate(string(doc))
	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGenerate_one_to_one_buy_CNPJ_and_sell_CPF_2(t *testing.T) {
	generatorApp := NewGeneratorApp()

	expected := `
<fragmento indice="CABECALHO" />
<br>
<u>TRANSMITENTE(S)</u>:> <span>MARIA EDUARDA DOS SANTOS, brasileiro, Solteiro (a), CPF nº 051.737.249-51, residente e domiciliado na Rua Nelson Carneiro Borges, nº 284, Bairro São Luiz, Brusque/SC.</span>.
<br/>
<u>ADQUIRENTE(S)</u>:><span>ABSOLUT CONSTRUTORA E INCORPORADORA LTDA, brasileiro, Solteiro (a), CPF nº 05.768.477/0001-35, residente e domiciliado na R CENTENARIO, nº 13, Bairro Santa Terezinha, Brusque/SC.</span>.
<br>
<u>FORMA DO TÍTULO</u>: Escritura Pública de Compra e Venda, lavrada pelo 1º Tabelionato de Notas e de Protesto de Brusque/SC, Livro 965, Folhas 124/126V, em 26/03/2025. 
<br/><u>VALOR</u>: R$ 210.000,00 (duzentos edez mil reais).
<br/><u>CONDIÇÕES</u>: Não constam.
<br/><u>OBSERVAÇÕES</u>:
<strong>ITBI</strong>: Recolhido no valor de R$ 4.600,00, com incidência sobre R$ 230.000,00, devidamente quitado. Nos termos do artigo 320 do CNCGFE/SC, o imóvel da presente matrícula, teve como valor atribuído de mercado, no
''quantum'' de R$ XXXXX. No ato da lavratura da Escritura Pública, foram apresentadas as certidões previstas em Lei. Com as demais cláusulas e condições da Escritura Pública. <strong> NO PRAZO REGULAMENTAR SERÁ EMITIDA A DOI</strong>.
<fragmento indice="FINALIZACAO_ATO" />.
`

	doc, err := os.ReadFile("../../infra/test_files/ato_consultar_tjsc_1_to_1_buy_and_sell_2.txt")
	require.NoError(t, err)

	got, err := generatorApp.Generate(string(doc))

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}
