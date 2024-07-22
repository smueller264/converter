package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type XMLReader struct{}

func (XMLReader) ReadFile(filePath string) (data Catalog, err error) {
	// Open our xmlFile
	xmlFile, err := os.Open("coffee_feed.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err := io.ReadAll(xmlFile)

	if err != nil {
		log.Fatal(err)
	}

	// we initialize our Catalog array
	var catalog Catalog
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &catalog)

	return catalog, err

}

func (XMLReader) CanHandle(format string) bool {
	if format == "xml" {
		return true
	} else {
		return false
	}
}
