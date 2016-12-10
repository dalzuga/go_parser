package main

import "encoding/xml"

// Author struct for XML
type Author struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Book struct for XML
type Book struct {
	ID      int      `xml:"id"`
	Title   string   `xml:"title"`
	Authors []Author `xml:"authors>author"`
}

// GoodReadsResult struct for XML
type GoodReadsResult struct {
	GoodreadsResponse xml.Name `xml:"GoodreadsResponse"`
	BookWrapper       []Book   `xml:"book"`
}
