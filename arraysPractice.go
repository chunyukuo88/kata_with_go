package main

import (
	"fmt"
)

func main() {
	products := []Product{
		Product{"The Wobbit", "123", 9.99},
		Product{"The Hobbit", "321", 1999.99},
	}

	thirdProduct := Product{"Lord of the Flies", "789", 5.36}
	updatedProducts := append(products, thirdProduct)
	fmt.Print(updatedProducts)
}

type Product struct {
	title string;
	id string;
	price float32;
}
