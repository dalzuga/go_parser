package main

import "encoding/xml"

// AuthorGRBQ struct for XML
type AuthorGRBQ struct {
	ID    int        `xml:"id"`
	Name  string     `xml:"name"`
	Books []BookGRBQ `xml:"books>book"`
}

// BookGRBQ struct for XML
type BookGRBQ struct {
	ID      int          `xml:"id"`
	Title   string       `xml:"title"`
	Authors []AuthorGRBQ `xml:"authors>author"`
}

// AuthorGRAQ struct for XML
type AuthorGRAQ struct {
	ID    int        `xml:"id"`
	Name  string     `xml:"name"`
	Books []BookGRAQ `xml:"books>book"`
}

// BookGRAQ struct for XML
type BookGRAQ struct {
	ID      int          `xml:"id"`
	Title   string       `xml:"title"`
	Authors []AuthorGRAQ `xml:"authors>author"`
}

// GoodReadsBookQuery struct for XML
type GoodReadsBookQuery struct {
	GoodreadsResponse xml.Name `xml:"GoodreadsResponse"`
	Book              BookGRBQ `xml:"book"`
}

// GoodReadsAuthorQuery struct for XML
type GoodReadsAuthorQuery struct {
	GoodreadsResponse xml.Name   `xml:"GoodreadsResponse"`
	Author            AuthorGRBQ `xml:"author"`
}
