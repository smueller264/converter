package parser

import "testing"

func TestReadFile(t *testing.T) {
	xmlReader := XMLReader{}
	_, err := xmlReader.ReadFile("../coffee_feed.xml")
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
