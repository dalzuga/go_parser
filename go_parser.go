package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

const (
	cOMMENTOPENEXP  = "<!--"
	cOMMENTCLOSEEXP = "-->"
	cDATAOPENEXP    = "<![CDATA["
	cDATACLOSEEXP   = "]]>"
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
		checkForCData(*scanner)
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

func checkForComment(scanner *bufio.Scanner) bool {
	CommentOpenExp := cOMMENTOPENEXP
	CommentCloseExp := cOMMENTCLOSEEXP

	lenCommentOpenExp := len(CommentOpenExp)
	flagOpenComment := false

	var i int8;

	for i = 0; i < lineLength; i++ {
		if lineLength-i < lenCommentOpenExp {
			break
		}
		if CommentOpenExp == line[i:i+lenCommentOpenExp] {
			flagOpenComment = true
			break
		}
	}

	lenCommentCloseExp := len(CommentCloseExp)
	flagCloseComment := false

	for ; i < lineLength; i++ {
		if lineLength-i < lenCommentOpenExp {
			break
		}
		if CommentOpenExp == line[i:i+lenCommentOpenExp] {
			flagOpenComment = true
			break
		}
	}

	return true
	// testing subStringInString function
	// if subStringInString("", "") {
	// 	fmt.Println("1true")
	// } else {
	// 	fmt.Println("1false")
	// }

	fmt.Println(CommentOpenExp + CommentCloseExp)
}

func checkForCData(scanner bufio.Scanner) {
	CDataOpenExp := cDATAOPENEXP
	CDataCloseExp := cDATACLOSEEXP

	line := scanner.Text()
	if subStringInString(CDataOpenExp, line) {
		fmt.Println("CData open found!")
		// fmt.Println("line:" + line)
	}

	if subStringInString(CDataCloseExp, line) {
		fmt.Println("CData close found!")
		// fmt.Println("line:" + line)
	}

	// fmt.Println(CDataCloseExp + CDataOpenExp)
}

func subStringInString(sub string, str string) bool {
	stringLength := len(str)
	subStringLength := len(sub)
	for i := 0; i < stringLength; i++ {
		if stringLength-i < subStringLength {
			return false
		}
		if sub == str[i:i+subStringLength] {
			return true
		}
	}
	return false
}

// Excepted means anything that is an exception in XML for not parsing the XML tags
// Includes CDATA and comments of type <!-- comment -->
// Returns nil if there are no changes
// It redundantly checks if the line needs to be scrubbed
// Returns true on success, false if an error occurred
func scrubExcepted(scanner *bufio.Scanner) {
	/* look for Excepted sections */
	line := scanner.Text()

	CDataOpenExp := cDATAOPENEXP
	CDataCloseExp := cDATACLOSEEXP
	CommentOpenExp := cOMMENTOPENEXP
	CommentCloseExp := cOMMENTCLOSEEXP

	lineLength := len(line)

	/* for loop */


	if flagOpenComment == false

	// unfinished code
	// for i := 0; i < lineLength; i++ {
	//
	// }

}
