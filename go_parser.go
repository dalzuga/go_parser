package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	f, err := os.Open("books.xml")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	captureTagRegExp := "</?(.+?)>" // captures the XML tag
	re := regexp.MustCompile(captureTagRegExp)

	var regExpResult []string

	mapXMLTags := make(map[string]bool)

	for scanner.Scan() {
		regExpResult = re.FindStringSubmatch(scanner.Text())
		if regExpResult != nil { // kind of wasteful but necessary check
			regExpResult = regExpResult[1:] // take out first element
		}

		for _, value := range regExpResult {
			if mapXMLTags[value] == false {
				mapXMLTags[value] = true
				fmt.Println(value)
			}
		}
	}
}