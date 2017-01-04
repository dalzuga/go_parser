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

/*
 * GRBQ stands for GoodReads Book Query
 * GRAQ stands for GoodReads Author Query
 */

func main() {
	AuthorID, err := getAuthorID("books.xml")
	if err != nil {
		log.Fatal(err)
	}

	mapTitles, err := requestAllBookTitles(AuthorID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mapTitles[0])
}

/*
 * getAuthorID takes a filename, loads into memory, and parses
 * the XML into the corresponding struct
 */

func getAuthorID(fileName string) (int, error) {
	fileBytes, err := ioutil.ReadFile("books.xml") // Read the GRBQ XML into memory
	if err != nil {
		return 0, err
	}

	var grbq GoodReadsBookQuery

	err = xml.Unmarshal(fileBytes, &grbq) // Parse the GBRQ XML to a GRBQ struct
	if err != nil {
		return 0, err
	}

	return grbq.Book.Authors[0].ID, nil // Return the Author ID from the GRBQ struct
}

func requestAllBookTitles(AuthorID int) (map[int]string, error) {
	fmt.Println(AuthorID)
	endpointBase := "https://www.goodreads.com/author/list.xml"
	page := 1
	var startBooks, endBooks, totalBooks int

	fmt.Println(startBooks, endBooks, totalBooks)

	/*
	 * Here, 'u' is a parsed object.
	 */
	u, err := url.Parse(endpointBase)
	if err != nil {
		return make(map[int]string), err
	}

	/*
	 * Here, the query 'q' is formulated
	 */
	q := u.Query()
	q.Set("key", `kDkKnUxiz8cRBJhVjrtSA`)
	q.Set("id", strconv.Itoa(AuthorID))
	q.Set("page", strconv.Itoa(page))

	/*
	 * Here, 's' is our full constructed URL string
	 */
	u.RawQuery = q.Encode()
	s := u.String()
	fmt.Println(s)

	req, err := http.NewRequest("GET", s, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	requestBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var graq GoodReadsAuthorQuery

	err = xml.Unmarshal(requestBytes, &graq)
	if err != nil {
		log.Fatal(err)
	}

	mapTitles := make(map[int]string)

	for key, bookValue := range graq.Author.Books.Book {
		mapTitles[key] = bookValue.Title
	}

	return mapTitles, nil
}
