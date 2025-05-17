package minuta

import (
	"fmt"
	"strings"
	"time"
)

const (
	DayMonthYearFormat = "02/01/2006"
)

const cabecalhoTemplate = `
<cabecalho>
  <b>{{TYPE_ATO}}.{{NUM_ATO}}-{{MATRICULA}}, </b> em {{DATA_REGISTRO}}.
  <br>
  <b>Prot. {{PROTOCOLO}},</b> datado de {{DATA_PROTOCOLO}}.
  <br>
  <u><b><maiusculo>{{TITLE_ATO}}</maiusculo></b></u>:
</cabecalho>
`

type CabecalhoParams struct {
	AtoType       string
	TitleAto      string
	NumAto        string
	Matricula     string
	RegistroDate  time.Time
	Protocolo     string
	ProtocoloDate time.Time
}

func registroDatef(date time.Time) string {
	return fmt.Sprintf("%d de %s de %d", date.Day(), date.Month(), date.Year())
}

func titleAtof(t string) string {
	return strings.ToUpper(t)
}

func Cabecalho(params CabecalhoParams) string {
	replacer := strings.NewReplacer(
		TypeAto.String(), params.AtoType,
		NumAto.String(), params.NumAto,
		Matricula.String(), params.Matricula,
		DataRegistro.String(), registroDatef(params.RegistroDate),
		Protocolo.String(), params.Protocolo,
		DataProtocolo.String(), params.ProtocoloDate.Format(DayMonthYearFormat),
		TitleAto.String(), titleAtof(params.TitleAto),
	)

	return replacer.Replace(cabecalhoTemplate)
}
