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

	var grbq GoodReadsBookQuery

	err = xml.Unmarshal(fileBytes, &grbq)

	AuthorID := grbq.Book.Authors[0].ID

	fmt.Println("AuthorID:", AuthorID)

	fileBytes, err = ioutil.ReadFile("authorlistbooks.xml")

	if err != nil {
		log.Fatal(err)
	}

	var graq GoodReadsAuthorQuery

	err = xml.Unmarshal(fileBytes, &graq)

	for _, bookValue := range graq.Author.Books.Book {
		fmt.Println(bookValue.Title)
	}

	fmt.Println("start: ", graq.Author.Books.Start)
	fmt.Println("end: ", graq.Author.Books.End)
	fmt.Println("total: ", graq.Author.Books.Total)

	fmt.Println("________________________________")

	fmt.Println(start, end, total, (end-start)/total)
}
