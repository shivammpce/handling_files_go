package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	data, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("error while opening file:", err)
		return
	}
	defer data.Close()

	val := csv.NewReader(data)
	val2, err := val.ReadAll()
	if err != nil {
		fmt.Println("error while reading file:", err)
		return
	}

	fileCreate, err := os.Create("res.csv")
	if err != nil {
		fmt.Println("error while creating file:", err)
		return
	}
	defer fileCreate.Close()

	csvFileWriter := csv.NewWriter(fileCreate)

	if err := csvFileWriter.WriteAll(val2); err != nil {
		fmt.Println("error while writing to res.csv:", err)
		return
	}

	time.Sleep(10 * time.Second)

	os.Remove("res.csv")
}
