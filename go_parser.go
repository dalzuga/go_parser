package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

/*
 * GRBQ stands for GoodReads Book Query
 * GRAQ stands for GoodReads Author Query
 */

func main() {
	var AuthorID int
	var err error

	argc := len(os.Args)

	if argc == 2 {
		AuthorID, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		AuthorID, err = getAuthorID("books.xml")
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(AuthorID)

	mapTitles, err := requestAllBookTitles(AuthorID)
	if err != nil {
		log.Fatal(err)
	}

	/* print the book titles */
	mapLength := len(mapTitles)
	// for i := 0; i < mapLength; i++ {
	// 	fmt.Println(i+1, mapTitles[i])
	// }
	fmt.Println("Map titles obtained:", mapLength)

	/* testing feature */
	max := 1
	for i := 116; i <= 1000; i++ {
		time.Sleep(time.Second)
		endpointBase := "https://www.goodreads.com/author/list.xml"
		page := 1
		_, more, err := requestPage(page, i, endpointBase)
		if err != nil {
			log.Fatal(err)
		}
		if more > 0 {
			fmt.Println("Current:", more)
			fmt.Println("AuthorID:", i)
		}
		if more > max {
			max = more
			fmt.Println("_________________MAX_________________")
			fmt.Println("Max:", max)
			fmt.Println("AuthorID:", i)
		}
	}
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

	mapTitles, more, err := requestPage(page, AuthorID, endpointBase)
	if err != nil {
		return make(map[int]string), err
	}

	if more > 0 {
		fmt.Println("There are more books in the API.")
		fmt.Println("Additional requests needed:", more)
	}

	// /* Make a map of channels, 1 channel per page */
	// channels := make(map[int]chan map[int]string)
	// for i := 2; i <= more+1; i++ {
	// 	channels[i] = make(chan map[int]string)
	// }
	//
	// /* Make 'more' (number of) requests */
	// for i := 2; i <= more+1; i++ {
	// 	// channelMaps := make(chan map[int]string)
	// 	fmt.Println("For loop. i =", i)
	// 	go func(i int) {
	// 		fmt.Println("Go func. i =", i)
	// 		moreTitles, _, err := requestPage(i, AuthorID, endpointBase)
	// 		if err != nil {
	// 			fmt.Println("This request failed:", i)
	// 			channels[i] <- make(map[int]string)
	// 		} else {
	// 			fmt.Println("Received: page", i)
	// 			fmt.Println("111")
	// 			channels[i] <- moreTitles
	// 			fmt.Println("222")
	// 		}
	// 	}(i)
	// }
	//
	// /* usually 30, but left variable in case API changes (untested) */
	// booksPerPage := len(mapTitles)
	//
	// /* Receive pages in order */
	// for i := 2; i <= more+1; i++ {
	// 	moreTitles := <-channels[i]
	// 	fmt.Println("Receiving channel:", i)
	//
	// 	/* add them to mapTitles (sequential) */
	// 	for j := 0; j <= booksPerPage; j++ {
	// 		if moreTitles[j] != "" {
	// 			mapTitles[(i-1)*booksPerPage+j] = moreTitles[j]
	// 		}
	// 	}
	// }

	return mapTitles, nil
}
