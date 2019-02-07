package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var book Books
	book.author = "Vikram"
	book.book_id = 1

	fmt.Println(book)

}
