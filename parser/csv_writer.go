package parser

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type CSVWriter struct{}

func (CSVWriter) WriteFile(filePath string, catalog Catalog) error {
	outputFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	err = gocsv.MarshalFile(catalog.Items, outputFile) // Use this to save the CSV back to the file

	return err

}
