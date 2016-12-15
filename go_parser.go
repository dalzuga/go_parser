package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

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

/* Here we are parsing the file into a struct */
func parseFile(fileName string, graq *GoodReadsAuthorQuery) {
	/*
	 * Here we are re-using fileBytes but for GRAQ, analogously
	 */

	var fileBytes []byte

	fileBytes, err := ioutil.ReadFile(fileName) // Read the GRAQ XML into memory
	if err != nil {
		log.Fatal(err)
	}

	err = xml.Unmarshal(fileBytes, &graq) // Parse the GRAQ XML to a GRBQ struct

	for _, bookValue := range graq.Author.Books.Book {
		fmt.Println(bookValue.Title)
	}
}

/*
 * GRBQ stands for GoodReads Book Query
 * GRAQ stands for GoodReads Author Query
 */

func main() {
	var AuthorID int
	getAuthorID("books.xml", &AuthorID)

	var graqPageOne GoodReadsAuthorQuery
	parseFile("authorlistbooks.xml", &graqPageOne)

	startBooks := graqPageOne.Author.Books.Start
	endBooks := graqPageOne.Author.Books.End
	totalBooks := graqPageOne.Author.Books.Total

	fmt.Println("_______________XML INFO_______________")
	fmt.Println(startBooks, endBooks, totalBooks, totalBooks/endBooks)

	/* 
         * Code below is for pagination.
	 * 
         * Here I am using a scope trick: I didn't know how to clear the contents of graq idiomatically,
         * so I declared another variable, graqOtherPages, inside the loop
	 */
	
	pageNumber := 1
	for totalBooks > endBooks {
		var graqOtherPages GoodReadsAuthorQuery
		fmt.Println("_______________________REQUEST________________________")
		makeHTTPRequest("https://www.goodreads.com/author/list.xml", AuthorID, pageNumber, &graqOtherPages)
		startBooks = graqOtherPages.Author.Books.Start
		endBooks = graqOtherPages.Author.Books.End
		totalBooks = graqOtherPages.Author.Books.Total
		for _, bookValue := range graqOtherPages.Author.Books.Book {
			fmt.Println(bookValue.Title)
		}
		pageNumber++
	}

	fmt.Println("Total requests:", pageNumber-1)
}

/*
 * makeHTTPRequest takes the full URL string, makes a request, and parses
 * the XML in the response into the struct pointed to by graq
 */
func makeHTTPRequest(uri string, AuthorID int, pageNumber int, graq *GoodReadsAuthorQuery) {

	client := &http.Client{}

	u, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Host:", u.Host)
	// u.Scheme = "https"
	// u.Host = "goodreads.com"

	q := u.Query()
	q.Set("key", `kDkKnUxiz8cRBJhVjrtSA`)
	q.Set("id", strconv.Itoa(AuthorID))
	q.Set("page", strconv.Itoa(pageNumber))

	// fmt.Println(q.Encode())

	u.RawQuery = q.Encode()

	// fmt.Println(u.Host)
	// fmt.Println(u.RequestURI())

	fullURL := u.String()
	fmt.Println(fullURL)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	requestBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = xml.Unmarshal(requestBytes, graq)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%#v", graq.Author.Books.Book[0].Title)
}
