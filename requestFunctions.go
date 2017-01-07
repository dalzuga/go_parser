package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

/* This function requests a page from the API. */
func requestPage(page int, AuthorID int, endpointBase string) (map[int]string, int, error) {
	req, err := prepareRequest(endpointBase, page, AuthorID)
	if err != nil {
		return make(map[int]string), 0, err
	}

	/*
	 * resp is of type *http.Response
	 */
	resp, err := doRequest(req)
	if err != nil {
		return make(map[int]string), 0, err
	}

	/*
	 * var 'more' is an int.
	 * If the API needs to paginate the response, more will indicate how many
	 * pages need to be requested for a full list of book titles.
	 * If there is no need to paginate, more will default to 0.
	 */
	mapTitles, more, err := parseResponse(resp)
	if err != nil {
		return make(map[int]string), 0, err
	}

	return mapTitles, more, nil
}

func prepareRequest(endpointBase string, page int, AuthorID int) (*http.Request, error) {
	/*
	 * Here, 'u' is a url object.
	 */
	u, err := url.Parse(endpointBase)
	if err != nil {
		return nil, err
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

	req, err := http.NewRequest("GET", s, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func doRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func parseResponse(resp *http.Response) (map[int]string, int, error) {
	requestBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return make(map[int]string), 0, err
	}

	var graq GoodReadsAuthorQuery

	err = xml.Unmarshal(requestBytes, &graq)
	if err != nil {
		return make(map[int]string), 0, err
	}

	var requestTitles = make(map[int]string)

	for key, bookValue := range graq.Author.Books.Book {
		requestTitles[key] = bookValue.Title
	}

	more, err := checkForMore(&graq)
	if err != nil {
		return make(map[int]string), 0, err
	}

	return requestTitles, more, nil
}

func checkForMore(graq *GoodReadsAuthorQuery) (int, error) {
	var start, end, total int

	start = graq.Author.Books.Start
	end = graq.Author.Books.End
	total = graq.Author.Books.Total

	if total != end {
		return (total-end)/(end-start+1) + 1, nil
	}

	return 0, nil
}
