package mini

import (
	"bookstore_system/db_access"
	"bookstore_system/model"
	"fmt"
)

func main() {

	bs := db_access.BookStore{}

	bs.AddBook(model.Book{ID: 1, Title: "The Hunger Games", Author: "Suzanne Colins", Quantity: 1, Availability: true})
	bs.AddBook(model.Book{ID: 2, Title: "Animal Farm", Author: "George Orwell", Quantity: 1, Availability: false})
	bs.AddBook(model.Book{ID: 3, Title: "Harry Potter", Author: "JK Rowling", Quantity: 1, Availability: true})

	book, err :=  bs.GetBook(3)
	if err != nil{
		fmt.Println("book not found", err)
	}else{
		fmt.Println(book)
	}

}