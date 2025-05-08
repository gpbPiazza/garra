package minuta

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate_one_to_one_buyer_CPF_and_seller_CNPJ(t *testing.T) {
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
<strong>ITBI</strong>: Recolhido no valor de R$ 2.900,00, com incidência sobre R$ 145.000,00, devidamente quitado. No ato da lavratura da Escritura Pública, foram apresentadas as certidões previstas em Lei. Com as demais cláusulas e condições da Escritura Pública. <strong> NO PRAZO REGULAMENTAR SERÁ EMITIDA A DOI</strong>.
<fragmento indice="FINALIZACAO_ATO" />.
`

	doc, err := os.ReadFile("../../infra/test_files/ato_consultar_tjsc_1_to_1_buyer_CPF_and_sellerr_CNPJ.txt")
	require.NoError(t, err)

	params := GenerateParams{
		DocStr:                      string(doc),
		IsTransmitenteOverqualified: false,
		IsAdquirenteOverqualified:   false,
	}

	got, err := generatorApp.Generate(params)
	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGenerate_one_to_one_buyer_CNPJ_and_sellerr_CPF(t *testing.T) {
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
<strong>ITBI</strong>: Recolhido no valor de R$ 4.600,00, com incidência sobre R$ 230.000,00, devidamente quitado. No ato da lavratura da Escritura Pública, foram apresentadas as certidões previstas em Lei. Com as demais cláusulas e condições da Escritura Pública. <strong> NO PRAZO REGULAMENTAR SERÁ EMITIDA A DOI</strong>.
<fragmento indice="FINALIZACAO_ATO" />.
`

	doc, err := os.ReadFile("../../infra/test_files/ato_consultar_tjsc_1_to_1_buyer_CNPJ_and_sellerr_CPF.txt")
	require.NoError(t, err)

	params := GenerateParams{
		DocStr:                      string(doc),
		IsTransmitenteOverqualified: false,
		IsAdquirenteOverqualified:   false,
	}

	got, err := generatorApp.Generate(params)

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGenerate_one_to_one_buyer_CNPJ_and_sellerr_CPF_2(t *testing.T) {
	generatorApp := NewGeneratorApp()

	expected := `
<fragmento indice="CABECALHO" />
<br>
<u>TRANSMITENTE(S)</u>:> <span>MARIA JOSE PEREIRA, supraqualificada.</span>.
<br/>
<u>ADQUIRENTE(S)</u>:><span>BBK EMPREENDIMENTOS IMOBILIÁRIOS LTDA., CNPJ nº 20.025.828/0001-01, com sede na rua RUA ALBERTO KLABUNDE, Nº 294 0, nº 294, Bairro ÁGUAS CLARAS, Brusque/SC.</span>.
<br>
<u>FORMA DO TÍTULO</u>: Escritura Pública de Compra e Venda, lavrada pelo 1º Tabelionato de Notas e de Protesto de Brusque/SC, Livro 965, Folhas 110/112V, em 26/03/2025. 
<br/><u>VALOR</u>: R$ 200.000,00 (duzentos mil reais).
<br/><u>CONDIÇÕES</u>: Não constam.
<br/><u>OBSERVAÇÕES</u>:
<strong>ITBI</strong>: Recolhido no valor de R$ 4.000,00, com incidência sobre R$ 200.000,00, devidamente quitado. No ato da lavratura da Escritura Pública, foram apresentadas as certidões previstas em Lei. Com as demais cláusulas e condições da Escritura Pública. <strong> NO PRAZO REGULAMENTAR SERÁ EMITIDA A DOI</strong>.
<fragmento indice="FINALIZACAO_ATO" />.
`

	// For this case we have a diffrente address for this company write down into escritura and the boxes from
	// tablionato register.
	// <u>ADQUIRENTE(S)</u>:><span>BBK EMPREENDIMENTOS IMOBILIARIOS LTDA., CNPJ nº 20.025.828/0001-01, com sede na Rua Sete de Setembro, nº 416, Sala 01, Bairro Santa Rita, Brusque/SC.</span>.

	doc, err := os.ReadFile("../../infra/test_files/ato_consultar_tjsc_1_to_1_buyer_CNPJ_and_sellerr_CPF_2.txt")
	require.NoError(t, err)

	params := GenerateParams{
		DocStr:                      string(doc),
		IsTransmitenteOverqualified: true,
		IsAdquirenteOverqualified:   false,
	}

	got, err := generatorApp.Generate(params)

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGenerate_one_to_one_buyer_CPF_and_sellerr_CNPJ_2(t *testing.T) {
	generatorApp := NewGeneratorApp()

	expected := `
<fragmento indice="CABECALHO" />
<br>
<u>TRANSMITENTE(S)</u>:> <span>Rmp;amp;R INCORPORAÇÃO LTDA, brasileiro, Separado Judicialmente (a), CPF nº 00.334.504/0001-48, residente e domiciliado na Avenida Primeiro de Maio, nº 346, Bairro Primeiro de Maio, São João do Itaperiú/SC.</span>.
<br/>
<u>ADQUIRENTE(S)</u>:><span>VILMAR PALOSCHI, brasileiro, Separado Judicialmente (a), CPF nº 548.480.839-15, residente e domiciliado na Rua Vice Prefeito Pedro Merizio, nº 399, Bairro Centro, São João do Itaperiú/SC.</span>.
<br>
<u>FORMA DO TÍTULO</u>: Escritura Pública de Compra e Venda, lavrada pelo 1º Tabelionato de Notas e de Protesto de Brusque/SC, Livro 965, Folhas 018/020V, em 20/03/2025. 
<br/><u>VALOR</u>: R$ 505.000,00 (quinhentos e cinco mil reais).
<br/><u>CONDIÇÕES</u>: Não constam.
<br/><u>OBSERVAÇÕES</u>:
<strong>ITBI</strong>: Recolhido no valor de R$ 11.800,00, com incidência sobre R$ 590.000,00, devidamente quitado. No ato da lavratura da Escritura Pública, foram apresentadas as certidões previstas em Lei. Com as demais cláusulas e condições da Escritura Pública. <strong> NO PRAZO REGULAMENTAR SERÁ EMITIDA A DOI</strong>.
<fragmento indice="FINALIZACAO_ATO" />.
`

	doc, err := os.ReadFile("../../infra/test_files/ato_consultar_tjsc_1_to_1_buyer_CPF_and_sellerr_CNPJ_2.txt")
	require.NoError(t, err)

	params := GenerateParams{
		DocStr:                      string(doc),
		IsTransmitenteOverqualified: false,
		IsAdquirenteOverqualified:   false,
	}

	got, err := generatorApp.Generate(params)

	require.NoError(t, err)
	assert.Equal(t, expected, got)
}
