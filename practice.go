package main

import "fmt"

func main() {
	var book1 BookData
	book1.ID = "123"
	book1.title = "The Hobbit"
	book1.description = "Book #1, a New York Times #1 Bestseller"
	book1.price = 12.03

	book2 := createBookInstance("123", "A Tree Grows in Brooklyn", "Some lame book", 9.13)

	book1.outputData(book1)
	book2.outputData(book2)
}

func createBookInstance(id string, title string, desc string, price float32) BookData {
	return BookData{id, title, desc, price}
}

type BookData struct {
	ID          string
	title       string
	description string
	price       float32
}

func (BookData) outputData(bookData BookData) {
	fmt.Printf("Here is the book: %v\n", bookData)
}

// Your Tasks
////////////////// 1) Create a new type / data structure for storing product data (e.g. about a book)
//////////////////		The data structure should contain these fields:
//////////////////		- ID
//////////////////		- Title / Name
//////////////////		- Short description
//////////////////		- price (number without currency, we'll just assume USD)
////////////////// 2) Create concrete instances of this data type in two different ways:
//////////////////		- Directly inside of the main function
//////////////////		- Inside of main, by using a "creation helper function"
//////////////////		(use any values for title etc. of your choice)
//////////////////		Output (print) the created data structures in the command line (in the main function)
////////////////// 3) Add a "connected function" that outputs the data + call that function from inside "main"
// 4) Change the program to fetch user input values for the different data fields
//		and create only one concrete instance with that entered data.
//		Output that instance data (via the connected function) then.
// 5) Bonus: Add a connected "store" function that writes that data into a file
//		The filename should be the unique id, the function should be called at the
//		end of main.
