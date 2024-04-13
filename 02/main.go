package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// open file
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	archiveFile, err := os.Create("archive.zip")
	if err != nil {
		panic(err)
	}

	defer archiveFile.Close()

	zipFileWriter := zip.NewWriter(archiveFile)

	fileCreate, err := zipFileWriter.Create("test.csv")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(fileCreate, f); err != nil {
		panic(err)
	}
	fmt.Println("closing zip archive...")
	zipFileWriter.Close()

	time.Sleep(10 * time.Second)
	fmt.Println("Waiting for 10 seconds...")

	removeZipFile := os.Remove("archive.zip")
	if removeZipFile != nil {
		log.Fatal(removeZipFile)
	}

}
