package main

import (
	"fmt"
	"strconv"
)

var book2 BookData

func main() {
	var book1 BookData
	book1.ID = "123"
	book1.title = "The Hobbit"
	book1.description = "Book #1, a New York Times #1 Bestseller"
	book1.price = 12.03

	book1.outputData(book1)

	book2 = createBookInstance()
}

func createBookInstance() BookData {
	fmt.Print("Please enter the book ID:\n> ")
	id, _ := reader.ReadString('\n')
	fmt.Print("Enter the title:\n> ")
	title, _ := reader.ReadString('\n')
	fmt.Print("Enter the description:\n> ")
	description, _ := reader.ReadString('\n')
	fmt.Print("Enter the price:\n> ")
	price, _ := reader.ReadString('\n')
	priceFloat, _ := strconv.ParseFloat(price, 64)
	return BookData{id, title, description, priceFloat}
}

type BookData struct {
	ID          string
	title       string
	description string
	price       float64
}

func (BookData) outputData(bookData BookData) {
	fmt.Printf("Here is the book: %v\n", bookData)
}

// func writeBookDataToFile(book BookData) {
// 	bookAsBytes := new(bytes.Buffer)
// 	json.NewEncoder(bookAsBytes).Encode(book)

// 	err := os.WriteFile("someBooks", bookAsBytes, 0666)
// 	check(err)
// }
