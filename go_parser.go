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

	var v GoodReadsBookQuery

	err = xml.Unmarshal(fileBytes, &v)

	AuthorID := v.Book.Authors[0].ID

	fmt.Println(AuthorID)

	fileBytes, err = ioutil.ReadFile("authorlistbooks.xml")

	if err != nil {
		log.Fatal(err)
	}

	var graq GoodReadsAuthorQuery

	err = xml.Unmarshal(fileBytes, &graq)

	firstBook := graq.Author.Books[0].Title

	fmt.Println(firstBook)

}
