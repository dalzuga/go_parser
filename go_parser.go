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

	/*
	 * This regular expression captures XML tags according to the XML spec.
	 * XML tags can start with '<' or '</'. Thus: '</?'
	 * The first character cannot be a punctuation character or a number,
	 * and it must not be empty. Thus: '[A-Za-z]{1}'.
	 * The '*?>' section is used for lazy matching until the closing '>'
	 * character.
	 * The parentheses '(', ')' are used for capturing and later extracting via
	 * the FindStringSubmatch function.
	 * The '[^\t\n\f\r ]' group matches anything that is not whitespace;
	 * together, with the placement of the closing parenthesis, this allows
	 * for not capturing attributes.
	 */

	captureTagRegExp := "</?([A-Za-z]{1}[^\t\n\f\r ]*?)>" // captures the XML tag
	re := regexp.MustCompile(captureTagRegExp)            // necessary syntax for extracting captures

	var regExpResult []string

	mapXMLTags := make(map[string]bool)

	for scanner.Scan() { // read each line
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

func checkForComments(scanner bufio.Scanner) {
	CommentOpenExp := "<!--"
	CommentCloseExp := "-->"

	fmt.Println(CommentOpenExp + CommentCloseExp)
}

func checkForCData(scanner bufio.Scanner) {
	CDataOpenExp := "<![CDATA["
	CDataCloseExp := "]]>"

	fmt.Println(CDataCloseExp + CDataOpenExp)
}
