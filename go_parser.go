package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Printf("hello, world\n")
	f, err := ioutil.ReadFile("books.xml")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", f)
}
