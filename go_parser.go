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

	type Author struct {
		ID   int    `xml:"id"`
		Name string `xml:"name"`
	}

	type Book struct {
		ID      int      `xml:"id"`
		Title   string   `xml:"title"`
		Authors []Author `xml:"authors>author"`
	}

	type Result struct {
		GoodreadsResponse xml.Name `xml:"GoodreadsResponse"`
		BookWrapper       []Book   `xml:"book"`
	}

	var v Result

	err = xml.Unmarshal(fileBytes, &v)

	fmt.Println(v.BookWrapper[0].Authors[0].Name)
}
