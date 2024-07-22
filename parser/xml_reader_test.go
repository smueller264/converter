package parser

import "testing"

func TestReadFile(t *testing.T) {
	xmlReader := XMLReader{}
	data := []byte{}
	_, err := xmlReader.MarshalData(data)
	if err != nil {
		t.Error(err)
	}
}

func TestCanHandle(t *testing.T) {
	xmlReader := XMLReader{}
	if xmlReader.CanHandle("xml") == false {
		t.Error("Wrong result")
	}
}
