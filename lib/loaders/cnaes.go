package loaders

import (
	"bytes"
	"encoding/csv"
	"gov-data/lib/pipeline"
	"gov-data/lib/storage"
	"io"
	"log"
)

type cnaeLoader struct {
	storage storage.Storage
}

func NewCnaeLoader(s storage.Storage) *cnaeLoader {
	return &cnaeLoader{
		storage: s,
	}
}

func (l *cnaeLoader) Load() error {
	// assumption: file exists at this location
	file := pipeline.Download("https://dadosabertos.rfb.gov.br/CNPJ/Cnaes.zip")
	// assumption: file is zipped
	unzippedFiles := pipeline.UnzipBytes(file)
	if len(unzippedFiles) != 1 {
		return ErrorDataSourceChanged
	}

	cnaesCsv := unzippedFiles[0]

	r := csv.NewReader(bytes.NewReader(cnaesCsv))
	r.Comma = ';'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error reading csv file")
		}

		label, err := pipeline.Win1252ToUtf8(record[1])
		if err != nil {
			log.Fatal("Error decoding data. Expected Win1252 encoded string, but got something else")
		}
		l.storage.Write(storage.Cnae{
			Code:  record[0],
			Label: label,
		})
	}

	l.storage.Close()
	return nil
}
