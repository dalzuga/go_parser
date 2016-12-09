package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileBytes, err := ioutil.ReadFile("books.xml") // Read file into memory

	if err != nil {
		log.Fatal(err)
	}

	var v Result

	err = xml.Unmarshal(fileBytes, &v)

	fmt.Println(v.BookWrapper[0].Authors[0].ID)
}
