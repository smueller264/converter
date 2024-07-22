package parser

import "testing"

func TestReadData(t *testing.T) {
	_, err := ReadData("test", "online")

	if err != nil {
		t.Error(err)
	}
}
