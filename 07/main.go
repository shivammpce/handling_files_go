package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	str := []string{"Café-society", "book", "pen"}

	Brother := []string{"Shubham", "Satyam"}
	Brother = append(Brother, "Rakesh")
	fmt.Printf("Brother's are : %s\n", Brother)

	for num, val := range Brother {
		fmt.Printf("list of brother's: %d %s\n", num, val)

	}
	fileCreate, err := os.Create("res.csv")
	if err != nil {
		fmt.Println("error while creating res.csv:", err)
		return
	}
	defer fileCreate.Close()

	csvFileWriter := csv.NewWriter(fileCreate)

	if err := csvFileWriter.Write(Brother); err != nil {
		fmt.Println("error while writing to res.csv:", err)
		return
	}

	if err := csvFileWriter.Write(str); err != nil {
		fmt.Println("error while writing to res.csv:", err)
		return
	}

	csvFileWriter.Flush()
}
