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
 * getAuthorID takes a filename, loads the file into memory, and parses the XML
 * into the corresponding struct to obtain the Goodreads author ID
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

/*
 * requestAllBookTitles retrieves all the book titles from the Goodreads API
 * @return: map[int]string
 */
func requestAllBookTitles(AuthorID int) (map[int]string, error) {
	endpointBase := "https://www.goodreads.com/author/list.xml"
	page := 1

	s, err := prepareRequest(endpointBase, page, AuthorID)
	if err != nil {
		return make(map[int]string), err
	}

	/*
	 * resp is of type *http.Response
	 */
	resp, err := doRequest(s)
	if err != nil {
		return make(map[int]string), err
	}

	/*
	 * If the API needs to paginate the response, set var 'more' to 'true'.
	 * Default is 'false'.
	 */
	mapTitles, more, err := parseResponse(resp)
	if err != nil {
		return make(map[int]string), err
	}

	if more {
		fmt.Println("There are more books in the API.")
	}

	return mapTitles, nil
}

func prepareRequest(endpointBase string, page int, AuthorID int) (string, error) {
	/*
	 * Here, 'u' is a url object.
	 */
	u, err := url.Parse(endpointBase)
	if err != nil {
		return "", err
	}

	/*
	 * Here, we make a query 'q' out of url object 'u'
	 */
	q := u.Query()
	q.Set("key", `kDkKnUxiz8cRBJhVjrtSA`)
	q.Set("id", strconv.Itoa(AuthorID))
	q.Set("page", strconv.Itoa(page))

	/*
	 * Here, 's' is our fully constructed URL string
	 */
	u.RawQuery = q.Encode()
	s := u.String()
	fmt.Println(s)

	return s, nil
}

func doRequest(s string) (*http.Response, error) {
	req, err := http.NewRequest("GET", s, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func parseResponse(resp *http.Response) (map[int]string, bool, error) {

	requestBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return make(map[int]string), false, err
	}

	var graq GoodReadsAuthorQuery

	err = xml.Unmarshal(requestBytes, &graq)
	if err != nil {
		return make(map[int]string), false, err
	}

	var requestTitles = make(map[int]string)

	for key, bookValue := range graq.Author.Books.Book {
		requestTitles[key] = bookValue.Title
	}

	more := checkForMore(&graq)

	return requestTitles, more, nil
}

func checkForMore(graq *GoodReadsAuthorQuery) bool {
	var end, total int

	end = graq.Author.Books.End
	total = graq.Author.Books.Total

	if total != end {
		return true
	}

	return false
}
