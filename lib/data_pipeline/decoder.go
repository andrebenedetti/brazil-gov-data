package data_pipeline

import "golang.org/x/text/encoding/charmap"

// For some reason, gov-br data comes in Windows1252 encoding.
// This function takes a Windows1252 string and decodes it to UTF-8.
// If you provide a string with any other encoding, the result will be incorrect.
func Win1252ToUtf8(str string) (string, error) {
	return charmap.Windows1252.NewDecoder().String(str)
}
