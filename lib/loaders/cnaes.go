package loaders

import (
	"bytes"
	"encoding/csv"
	"gov-data/lib/pipeline"
)

// CNAE stands for Classificação Nacional de Atividades Econômicas, or
// National Classification of Economic Activities.
// Each entry is a numeric code and a string that identifies an economic activity.
// For example, this CNAE identifies nursery activities:
// "8650001";"Atividades de enfermagem"

// We store codes as strings to handle CNAES in the form "0112199"
type Cnae struct {
	Code  string `json:"code"`
	Label string `json:"label"`
}

type CnaeLoader struct {
}

func (l *CnaeLoader) Load() ([]Cnae, error) {
	// assumption: file exists at this location
	file := pipeline.Download("https://dadosabertos.rfb.gov.br/CNPJ/Cnaes.zip")
	// assumption: file is zipped
	unzippedFiles := pipeline.UnzipBytes(file)
	if len(unzippedFiles) != 1 {
		return []Cnae{}, ErrorDataSourceChanged
	}

	cnaesCsv := unzippedFiles[0]

	r := csv.NewReader(bytes.NewReader(cnaesCsv))
	r.Comma = ';'
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	cnaes := make([]Cnae, len(records))
	for i, record := range records {
		label, err := pipeline.Win1252ToUtf8(record[1])
		if err != nil {
			return cnaes, ErrorDataSourceChanged
		}

		cnaes[i] = Cnae{
			Code:  record[0],
			Label: label,
		}
	}

	return cnaes, nil
}
