package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the input files "data1.txt" and "data2.txt"
	data1, err := os.Open("data1.txt")
	if err != nil {
		fmt.Println("error while opening data1.txt:", err)
		return
	}
	defer data1.Close()

	data2, err := os.Open("data2.txt")
	if err != nil {
		fmt.Println("error while opening data2.txt:", err)
		return
	}
	defer data2.Close()

	// Read the CSV data from "data1.txt" and "data2.txt"
	val1 := csv.NewReader(data1)
	val1Data, err := val1.ReadAll()
	if err != nil {
		fmt.Println("error while reading data1.txt:", err)
		return
	}

	val2 := csv.NewReader(data2)
	val2Data, err := val2.ReadAll()
	if err != nil {
		fmt.Println("error while reading data2.txt:", err)
		return
	}

	// Combine the data from both files
	combinedData := append(val1Data, val2Data...)

	// Create the output CSV file "res.csv"
	fileCreate, err := os.Create("res.csv")
	if err != nil {
		fmt.Println("error while creating res.csv:", err)
		return
	}
	defer fileCreate.Close()

	// Initialize the CSV writer
	csvFileWriter := csv.NewWriter(fileCreate)

	// Write the combined CSV data to "res.csv"
	if err := csvFileWriter.WriteAll(combinedData); err != nil {
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
