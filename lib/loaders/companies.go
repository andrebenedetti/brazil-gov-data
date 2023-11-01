package loaders

import (
	"bytes"
	"encoding/csv"
	"gov-data/lib/pipeline"
	"gov-data/lib/storage"
	"io"
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

type companyLoader struct {
	storage storage.Storage
}

func NewCompanyLoader(s storage.Storage) companyLoader {
	return companyLoader{
		storage: s,
	}
}

func (l *companyLoader) Load() error {
	file := pipeline.Download("https://dadosabertos.rfb.gov.br/CNPJ/Empresas0.zip")
	unzippedFiles := pipeline.UnzipBytes(file)
	if len(unzippedFiles) != 1 {
		return ErrorDataSourceChanged
	}

	companiesCsv := unzippedFiles[0]

	r := csv.NewReader(bytes.NewReader(companiesCsv))
	r.Comma = ';'

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		l.storage.Write(Company{
			BasicCnpj:                        record[0],
			Name:                             record[1],
			LegalNature:                      record[2],
			LegalRepresentativeQualification: record[3],
			NominalCapital:                   record[4],
			Size:                             record[5],
			FederalEntityInCharge:            record[6],
		})
	}

	return nil
}
