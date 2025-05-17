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
<u>TRANSMITENTE(S)</u>:> <span><strong>CONSTRUARTE CONSTRUTORA E INCORPORADORA LTDA.</strong>, CNPJ nº 29.235.977/0001-70, com sede na rua Rua Luiz Eccel, nº 1, Bairro Paquetá, Brusque/SC.</span>
<br/>
<u>ADQUIRENTE(S)</u>:><span><strong>SIDNEI ANTONIO GATTIS</strong>, brasileiro, solteiro, CPF nº 037.561.669-10, residente e domiciliado na Rua Azambuja, nº 541, Bairro Azambuja, Brusque/SC.</span>
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
<u>TRANSMITENTE(S)</u>:> <span><strong>MARIA EDUARDA DOS SANTOS</strong>, brasileiro, solteira, CPF nº 051.737.249-51, residente e domiciliado na Rua Nelson Carneiro Borges, nº 284, Bairro São Luiz, Brusque/SC.</span>
<br/>
<u>ADQUIRENTE(S)</u>:><span><strong>ABSOLUT CONSTRUTORA E INCORPORADORA LTDA.</strong>, CNPJ nº 05.768.477/0001-35, com sede na rua R CENTENARIO, nº 13, Bairro Santa Terezinha, Brusque/SC.</span>
<br>
<u>FORMA DO TÍTULO</u>: Escritura Pública de Compra e Venda, lavrada pelo 1º Tabelionato de Notas e de Protesto de Brusque/SC, Livro 965, Folhas 124/126V, em 26/03/2025. 
<br/><u>VALOR</u>: R$ 210.000,00 (duzentos e dez mil reais).
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
<u>TRANSMITENTE(S)</u>:> <span><strong>MARIA JOSE PEREIRA</strong>, supraqualificada.</span>
<br/>
<u>ADQUIRENTE(S)</u>:><span><strong>BBK EMPREENDIMENTOS IMOBILIÁRIOS LTDA.</strong>, CNPJ nº 20.025.828/0001-01, com sede na rua RUA ALBERTO KLABUNDE, Nº 294 0, nº 294, Bairro ÁGUAS CLARAS, Brusque/SC.</span>
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
<u>TRANSMITENTE(S)</u>:> <span><strong>RMP;AMP;R INCORPORAÇÃO LTDA.</strong>, CNPJ nº 00.334.504/0001-48, com sede na rua Avenida Primeiro de Maio, nº 346, Bairro Primeiro de Maio, São João do Itaperiú/SC.</span>
<br/>
<u>ADQUIRENTE(S)</u>:><span><strong>VILMAR PALOSCHI</strong>, brasileiro, divorciado, CPF nº 548.480.839-15, residente e domiciliado na Rua Vice Prefeito Pedro Merizio, nº 399, Bairro Centro, São João do Itaperiú/SC.</span>
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

func TestGenerate_one_to_one_offset_of_some_key_in_between_pages_bug_1(t *testing.T) {
	generatorApp := NewGeneratorApp()

	expected := `
<fragmento indice="CABECALHO" />
<br>
<u>TRANSMITENTE(S)</u>:> <span><strong>MORATTA EMPREENDIMENTOS IMOBILIÁRIOS LTDA.</strong>, CNPJ nº 08.475.810/0001-06, com sede na rua RUA MATHILDE SCHAEFER, nº 173, Bairro SÃO LUIZ, Brusque/SC.</span>
<br/>
<u>ADQUIRENTE(S)</u>:><span><strong>NADIR HASSMANN</strong>, brasileiro, divorciada, CPF nº 823.199.959-00, residente e domiciliado na RUA OTTO KRIEGER, nº 40, Bairro SÃO LUIZ, Brusque/SC.</span>
<br>
<u>FORMA DO TÍTULO</u>: Escritura Pública de Compra e Venda, lavrada pelo 1º Tabelionato de Notas e de Protesto de Brusque/SC, Livro 967, Folhas 018/020V, em 14/04/2025. 
<br/><u>VALOR</u>: R$ 340.000,00 (trezentos e quarenta mil reais).
<br/><u>CONDIÇÕES</u>: Não constam.
<br/><u>OBSERVAÇÕES</u>:
<strong>ITBI</strong>: Recolhido no valor de R$ 6.800,00, com incidência sobre R$ 340.000,00, devidamente quitado. No ato da lavratura da Escritura Pública, foram apresentadas as certidões previstas em Lei. Com as demais cláusulas e condições da Escritura Pública. <strong> NO PRAZO REGULAMENTAR SERÁ EMITIDA A DOI</strong>.
<fragmento indice="FINALIZACAO_ATO" />.
`

	doc, err := os.ReadFile("../../infra/test_files/offset_of_some_key_in_between_pages_bug.txt")
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

func TestGenerate_one_to_one_case_to_many_start_keys_bug_2(t *testing.T) {
	generatorApp := NewGeneratorApp()

	expected := `
<fragmento indice="CABECALHO" />
<br>
<u>TRANSMITENTE(S)</u>:> <span><strong>FABRICIA RIBEIRO DOS SANTOS</strong>, brasileiro, casada, CPF nº 902.091.845-15, residente e domiciliado na Dom Joaquim, nº 155, Bairro Cedrinho, Brusque/SC.</span>
<br/>
<u>ADQUIRENTE(S)</u>:><span><strong>JOÃO PAULO MENDONÇA MELO</strong>, brasileiro, solteiro, CPF nº 075.979.619-01, residente e domiciliado na Rua Francisco Debatin, nº 22, Bairro Águas Claras, Brusque/SC.</span>
<br>
<u>FORMA DO TÍTULO</u>: Escritura Pública de Compra e Venda, lavrada pelo 1º Tabelionato de Notas e de Protesto de Brusque/SC, Livro 968, Folhas 095/097V, em 02/05/2025. 
<br/><u>VALOR</u>: R$ 180.000,00 (cento e oitenta mil reais).
<br/><u>CONDIÇÕES</u>: Não constam.
<br/><u>OBSERVAÇÕES</u>:
<strong>ITBI</strong>: Recolhido no valor de R$ 3.600,00, com incidência sobre R$ 180.000,00, devidamente quitado. No ato da lavratura da Escritura Pública, foram apresentadas as certidões previstas em Lei. Com as demais cláusulas e condições da Escritura Pública. <strong> NO PRAZO REGULAMENTAR SERÁ EMITIDA A DOI</strong>.
<fragmento indice="FINALIZACAO_ATO" />.
`

	doc, err := os.ReadFile("../../infra/test_files/test_bug_2.txt")
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
