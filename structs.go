package main

import "encoding/xml"

// Author struct for XML
type Author struct {
	ID    int    `xml:"id"`
	Name  string `xml:"name"`
	Books []Book `xml:"books>book"`
}

// Book struct for XML
type Book struct {
	ID      int      `xml:"id"`
	Title   string   `xml:"title"`
	Authors []Author `xml:"authors>author"`
}

// GoodReadsBookQuery struct for XML
type GoodReadsBookQuery struct {
	GoodreadsResponse xml.Name `xml:"GoodreadsResponse"`
	Book              Book     `xml:"book"`
}

// GoodReadsAuthorQuery struct for XML
type GoodReadsAuthorQuery struct {
	GoodreadsResponse xml.Name `xml:"GoodreadsResponse"`
	Author            Author   `xml:"author"`
}
