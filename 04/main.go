package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the input file "data.txt"
	data, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("error while opening file:", err)
		return
	}
	defer data.Close()

	// Read the CSV data from "data.txt"
	val := csv.NewReader(data)
	val2, err := val.ReadAll()
	if err != nil {
		fmt.Println("error while reading file:", err)
		return
	}

	// Create the output CSV file "res.csv"
	fileCreate, err := os.Create("res.csv")
	if err != nil {
		fmt.Println("error while creating file:", err)
		return
	}
	defer fileCreate.Close()

	// Initialize the CSV writer
	csvFileWriter := csv.NewWriter(fileCreate)

	// Write the CSV data to "res.csv"
	if err := csvFileWriter.WriteAll(val2); err != nil {
		fmt.Println("error while writing to res.csv:", err)
		return
	}

	// Flush the data to the file and check for errors
	csvFileWriter.Flush()
	if err := csvFileWriter.Error(); err != nil {
		fmt.Println("error while flushing to res.csv:", err)
		return
	}

	fmt.Println("CSV data has been successfully written to res.csv")
}
