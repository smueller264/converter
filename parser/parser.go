package parser

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"os"
)

type Item struct {
	XMLName       xml.Name `xml:"item" csv:"item"`
	Entity_id     string   `xml:"entity_id" csv:"entity_id"`
	Category_name string   `xml:"CategoryName" csv:"CategoryName"`
	Sku           string   `xml:"sku" csv:"sku"`
	Name          string   `xml:"name" cvs:"name"`
	Description   string   `xml:"description" csv:"description"`
	Shortdesc     string   `xml:"shortdesc" csv:"shortdesc"`
	Price         string   `xml:"price" csv:"price"`
	Link          string   `xml:"link" csv:"link"`
	Image         string   `xml:"image" csv:"image"`
	Brand         string   `xml:"Brand" csv:"brand"`
	Rating        string   `xml:"Rating" csv:"rating"`
	Caffeine_type string   `xml:"CaffeineType" csv:"CaffeineType"`
	Count         string   `xml:"Count" csv:"Count"`
	Flavored      string   `xml:"Flavored" csv:"Flavored"`
	Seasonal      string   `xml:"Seasonal" csv:"Seasonal"`
	Instock       string   `xml:"Instock" csv:"Instock"`
	Facebook      string   `xml:"Facebook" csv:"Facebook"`
	Is_k_cup      string   `xml:"IsKCup" csv:"IsKCup"`
}

type Catalog struct {
	XMLName xml.Name `xml:"catalog"`
	Items   []Item   `xml:"item"`
}

type Reader interface {
	MarshalData(incdata []byte) (catalog Catalog, err error)
	CanHandle(format string) bool
}

type Writer interface {
	WriteFile(catalog Catalog) error
}

// Here you can add new Readers to extend functionality
var xmlReader = XMLReader{}

var readers = []Reader{
	xmlReader,
}

// Start the parsing process
func Parse(fileOrigin string, filePath string) {
	//reading the file
	data, err := ReadData(fileOrigin, filePath)

	if err != nil {
		log.Fatal(err)
	}

	//saving the file, here you can add any ouput format by adding a different Writer
	csvWriter := CSVWriter{}
	csvWriter.WriteFile("output.csv", *data)

}

// Reading the data
func ReadData(fileOrigin string, filepath string) (*Catalog, error) {

	var incdata []byte
	var err error

	//get online file
	if fileOrigin == "online" {
		resp, err := http.Get(filepath)
		if err != nil {
			log.Fatal("Error in getting file online")
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatal("Error in Status code")
		}

		incdata, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error in reading online file")
		}
		//get offline file
	} else {
		// Open our xmlFile
		file, err := os.Open(filepath)

		// if we os.Open returns an error then handle it
		if err != nil {
			log.Fatal(err)
		}

		// defer the closing of our xmlFile so that we can parse it later on
		defer file.Close()

		// read our opened xmlFile as a byte array.
		incdata, err = io.ReadAll(file)

		if err != nil {
			log.Fatal(err)
		}
	}

	marshaledData := Catalog{}

	for _, reader := range readers {
		if reader.CanHandle("xml") {
			marshaledData, err = reader.MarshalData(incdata)
			if err != nil {
				return nil, err
			}

		}
	}
	return &marshaledData, nil

}
