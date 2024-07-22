package main

import (
	"log"
	"os"

	"github.com/smueller264/converter/parser"
)

func main() {

	csvWriter := parser.CSVWriter{}

	//Setting up logging to logfile
	file, err := os.OpenFile("errorlog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	if err != nil {
		log.Fatal(err)
	}

	//reading the file
	data, err := parser.ReadData()

	if err != nil {
		log.Fatal(err)
	}

	//saving the file
	csvWriter.WriteFile("output.csv", *data)

}
