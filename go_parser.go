package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
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

	startBooks := graq.Author.Books.Start
	endBooks := graq.Author.Books.End
	totalBooks := graq.Author.Books.Total

	fmt.Println(startBooks, endBooks, totalBooks, totalBooks/endBooks)

	/* Code below is for pagination, need to code makeHTTPRequest */
	// for totalBooks < endBooks {
	// 	makeHTTPRequest(url string, AuthorID, &graq)
	// 	for _, bookValue := range graq.Author.Books.Book {
	// 		fmt.Println(bookValue.Title)
	// 	}
	// }

	makeHTTPRequest("https://www.goodreads.com/author/list.xml", AuthorID, &graq)
}

func makeHTTPRequest(uri string, AuthorID int, graq *GoodReadsAuthorQuery) error {

	// client := &http.Client{}

	u, err := url.Parse(uri)

	q := u.Query()
	q.Set("key", `kDkKnUxiz8cRBJhVjrtSA`)
	q.Set("id", `4`)

	fmt.Println(q.Encode())

	// resp, err := client.Do(req)
	//
	// fmt.Println(resp.Body)

	return err
}
