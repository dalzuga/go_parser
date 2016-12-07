package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("hello, world\n")
	f, err := os.Open("books.xml")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
