package extractor

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGremio(t *testing.T) {
	doc, err := os.ReadFile("../../infra/test_files/ato_consultar_tjsc_many_buyer_to_one_seller.txt")
	require.NoError(t, err)

	text, err := cutOffSet(string(doc), "Outorgado", "EscrituraAssinada na Serventia")
	require.NoError(t, err)

	// partes := strings.Split(text, "Parte")

	partes := strings.SplitAfter(text, "Parte")
	require.Len(t, partes, 3)

	// partes := strings.(text, "Parte")

	for _, p := range partes {
		fmt.Println(p)
		fmt.Print("\n")
	}
}

func TestGremio_1_1(t *testing.T) {
	doc, err := os.ReadFile("../../infra/test_files/ato_consultar_tjsc_1_to_1_buyer_CNPJ_and_sellerr_CPF_2.txt")
	require.NoError(t, err)

	text, err := cutOffSet(string(doc), "Outorgado", "EscrituraAssinada na Serventia")
	fmt.Println(text)
	require.NoError(t, err)

	partes := strings.SplitAfter(text, "Parte")
	require.Len(t, partes, 2)

	// partes := strings.(text, "Parte")

	for _, p := range partes {
		fmt.Println(p)
		fmt.Print("\n")
	}
}
