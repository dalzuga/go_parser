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

	for scanner.Scan() {
		// open, optionally '/', capture (xmltag non-greedy), close
		regExpResult = re.FindStringSubmatch(scanner.Text())
		if regExpResult != nil {
			fmt.Println(regExpResult[1])
		}
	}
}
