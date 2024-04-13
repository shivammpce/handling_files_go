package main

import (
	"fmt"
	"os"
)

func main() {
	m := map[string]int{
		"Shivam": 1,
		"Anupam": 2,
		"Safwan": 3,
		"Sangam": 4,
	}

	for key, val := range m {
		fmt.Println(key, val)
	}

	file, err := os.Create("data.csv")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("file: %v\n", file)

}
