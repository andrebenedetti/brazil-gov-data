package loaders

import (
	"bytes"
	"encoding/csv"
	"fmt"
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

func (l *CompanyLoader) Load() error {
	file := pipeline.Download("https://dadosabertos.rfb.gov.br/CNPJ/Empresas0.zip")
	unzippedFiles := pipeline.UnzipBytes(file)
	if len(unzippedFiles) != 1 {
		return ErrorDataSourceChanged
	}

	companiesCsv := unzippedFiles[0]

	r := csv.NewReader(bytes.NewReader(companiesCsv))
	r.Comma = ';'
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	// WIP
	companies := make([]Company, len(records))
	for i, record := range records {
		if i%20 == 0 {
			PrintMemUsage()
			fmt.Println(record)
		}

		// We can't parse all the records and return it as a slice as we were doing before, as it will end up
		// occupying way too much space.
		// We need to change Storage's interface to support writing bytes into it, chunks by chunks
		// Then, we should pass a storage into the Loader
		// Company{
		// 	BasicCnpj:                        record[0],
		// 	Name:                             record[1],
		// 	LegalNature:                      record[2],
		// 	LegalRepresentativeQualification: record[3],
		// 	NominalCapital:                   record[4],
		// 	Size:                             record[5],
		// 	FederalEntityInCharge:            record[6],
		// }
	}

	return nil
}
