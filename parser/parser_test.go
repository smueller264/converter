package parser

import "testing"

func TestReadData(t *testing.T) {
	_, err := ReadData()

	if err != nil {
		t.Error(err)
	}
}
