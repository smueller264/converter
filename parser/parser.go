package parser

import (
	"encoding/xml"
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
	ReadFile(filepath string) (catalog Catalog, err error)
	CanHandle(format string) bool
}

type Writer interface {
	WriteFile(catalog Catalog) error
}

type ReaderController struct {
}

func ReadData() (*Catalog, error) {
	xmlReader := XMLReader{}
	readers := []Reader{xmlReader}

	var err error
	data := Catalog{}

	for _, reader := range readers {
		if reader.CanHandle("xml") {
			data, err = reader.ReadFile("/..coffee_feed.xml")
			if err != nil {
				return nil, err
			}

		}
	}
	return &data, nil

}
