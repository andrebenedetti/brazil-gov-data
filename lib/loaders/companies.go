package loaders

import (
	"bytes"
	"encoding/csv"
	"gov-data/lib/pipeline"
)

// BasicCnpj is the first 8 digits of the CNPJ, the brazilian document that
// identifies a company. 
type Company struct {
	BasicCnpj                        string `json:"basicCnpj"`
	Name                             string `json:"name"`
	LegalNature                      string `json:"legalNature"`
	LegalRepresentativeQualification string `json:"legalRepresentativeQualification"`
	NominalCapital                   string `json:"nominalCapital"`
	Size                             string `json:"size"`
	FederalEntityInCharge            string `json:"federalEntityInCharge"`
}

type CompanyLoader struct {
}

// TODO: Need a better strategy to use less working memory
func (l *CompanyLoader) Load() ([]Company, error) {
	file := pipeline.Download("https://dadosabertos.rfb.gov.br/CNPJ/Empresas0.zip")
	unzippedFiles := pipeline.UnzipBytes(file)
	if len(unzippedFiles) != 1 {
		return []Company{}, ErrorDataSourceChanged
	}

	companiesCsv := unzippedFiles[0]

	r := csv.NewReader(bytes.NewReader(companiesCsv))
	r.Comma = ';'
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	companies := make([]Company, len(records))
	for i, record := range records {
		companies[i] = Company{
			BasicCnpj:                        record[0],
			Name:                             record[1],
			LegalNature:                      record[2],
			LegalRepresentativeQualification: record[3],
			NominalCapital:                   record[4],
			Size:                             record[5],
			FederalEntityInCharge:            record[6],
		}
	}

	return companies, nil
}
