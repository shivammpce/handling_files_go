package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []Server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type Server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, Server{"Shivam", "127.0.0.1"})
	v.Svs = append(v.Svs, Server{"Singh", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, " ", " ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
