package main

import "fmt"

func main() {
	var books []string

	books = append(books, "Buku 1")
	books = append(books, "Buku 2")
	books = append(books, "Buku 3")

	for _, book := range books {
		fmt.Println(book)
	}

	shops := []string{
		"Toko 1",
		"Toko 2",
	}

	for index, shop := range shops {
		fmt.Printf("Index ke-%d: %s", index, shop)
	}
}
