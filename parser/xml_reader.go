package parser

import (
	"encoding/xml"
)

type XMLReader struct{}

func (XMLReader) MarshalData(incdata []byte) (data Catalog, err error) {

	// we initialize our Catalog array
	var catalog Catalog
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(incdata, &catalog)

	return catalog, err

}

func (XMLReader) CanHandle(format string) bool {
	if format == "xml" {
		return true
	} else {
		return false
	}
}
