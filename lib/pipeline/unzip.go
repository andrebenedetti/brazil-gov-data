package pipeline

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
)

func UnzipBytes(z []byte) [][]byte {
	zipReader, err := zip.NewReader(bytes.NewReader(z), int64(len(z)))
	if err != nil {
		log.Fatal(err)
	}

	unzippedFiles := make([][]byte, 0)
	for _, zipFile := range zipReader.File {
		unzippedFileBytes, err := readZipFile(zipFile)
		if err != nil {
			log.Println(err)
			continue
		}

		unzippedFiles = append(unzippedFiles, unzippedFileBytes)
	}

	return unzippedFiles
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}
