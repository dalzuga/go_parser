package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

/*
 * GRBQ stands for GoodReads Book Query
 * GRAQ stands for GoodReads Author Query
 */

func main() {
	var AuthorID int
	getAuthorID("books.xml", &AuthorID)

	bookTitles, err := requestAllBookTitles(AuthorID)
	fmt.Println(reflect.TypeOf(bookTitles))
	if err != nil {
		log.Fatal(err)
	}
}

/*
 * getAuthorID takes a filename, loads into memory, and parses
 * the XML into the corresponding struct
 */

func getAuthorID(fileName string, AuthorID *int) {
	fileBytes, err := ioutil.ReadFile("books.xml") // Read the GRBQ XML into memory
	if err != nil {
		log.Fatal(err)
	}

	var grbq GoodReadsBookQuery

	err = xml.Unmarshal(fileBytes, &grbq) // Parse the GBRQ XML to a GRBQ struct
	if err != nil {
		log.Fatal(err)
	}

	*AuthorID = grbq.Book.Authors[0].ID // Get the Author ID from the GRBQ struct
}

func requestAllBookTitles(AuthorID int) (map[string]int, error) {
	fmt.Println(AuthorID)

	return make(map[string]int), nil
}
