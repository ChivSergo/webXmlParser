package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/ChivSergo/webXmlParser.git/xmlforparse"

	"golang.org/x/net/html/charset"
)

func main() {
	file, err := os.Open("mprocessors.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	in := bufio.NewReader(file)

	var compoents xmlforparse.MComponents
	decoder := xml.NewDecoder(in)
	decoder.CharsetReader = charset.NewReaderLabel

	decoder.Decode(&compoents)

	for _, processor := range compoents.Processors {
		fmt.Printf("processor Name: %s \n", processor.Name)
	}

}
