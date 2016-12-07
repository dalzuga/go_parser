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
	for scanner.Scan() {
		selector := "book"
		matched, err := regexp.MatchString("</?"+selector+">", scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if matched {
			fmt.Println(scanner.Text()) // Println will add back the final '\n'
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
