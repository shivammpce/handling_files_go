package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type ShoppingRecord struct {
	Vegetable string
	Fruit     string
}

func createShoppingList(data [][]string) []ShoppingRecord {
	var shoppingList []ShoppingRecord
	for i, line := range data {
		if i > 0 { // omit header line
			var rec ShoppingRecord
			for j, field := range line {
				if j == 0 {
					rec.Vegetable = field
				} else if j == 1 {
					rec.Fruit = field
				}
			}
			shoppingList = append(shoppingList, rec)
		}
	}
	return shoppingList
}

func main() {
	// open file
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	shoppingList := createShoppingList(data)

	// print the array
	fmt.Printf("%+v\n", shoppingList)
}
