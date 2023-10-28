package data_loaders

import "gov-data/lib/downloader"

// CNAE stands for Classificação Nacional de Atividades Econômicas, or
// National Classification of Economic Activities.
// Each entry is a numeric code and a string that identifies an economic activity.
// For example, this CNAE identifies nursery activities:
// "8650001";"Atividades de enfermagem"

// We store codes as strings to handle CNAES in the form "0112199"
type Cnae struct {
	code  string
	label string
}

type CnaeLoader struct {
}

// Download Raw does not do any type of parsing over the file.
// Just returns it as it is.
func (l *CnaeLoader) DownloadRaw() []byte {
	url := "https://dadosabertos.rfb.gov.br/CNPJ"
	filename := "Cnaes.zip"

	return downloader.Download(url + filename)
}

// func (l *CnaeLoader) Parse(file []byte) []Cnae {

// }
