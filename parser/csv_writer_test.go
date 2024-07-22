package parser

import "testing"

func TestWriteFile(t *testing.T) {
	csvWriter := CSVWriter{}
	catalog := Catalog{}

	err := csvWriter.WriteFile("../output.csv", catalog)

	if err != nil {
		t.Error(err)
	}
}
