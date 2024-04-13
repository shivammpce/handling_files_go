package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Note struct {
	XMLName xml.Name `xml:"note"`
	To      string   `xml:"to"`
	From    string   `xml:"from"`
	Heading string   `xml:"heading"`
	Body    string   `xml:"body"`
}

// type Recurlyservers struct {
// 	XMLName     xml.Name `xml:"servers"`
// 	Version     string   `xml:"version,attr"`
// 	Svs         []Server `xml:"server"`
// 	Description string   `xml:",innerxml"`
// }

// type Server struct {
// 	ServerName string `xml:"serverName"`
// 	ServerIP   string `xml:"serverIP"`
// }

func main() {
	file, err := os.Open("server.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	v := Note{}

	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("%+v\n", v)
}
