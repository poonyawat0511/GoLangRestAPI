package models

// Book struct to hold book data
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Books is the in-memory data for books
var Books []Book = []Book{
	{ID: 1, Title: "1984", Author: "George Orwell"},
	{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
}
