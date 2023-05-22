package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	counter := make(chan int, 1)
	counter <- 0

	var books []string

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			currentCounter := <-counter
			currentCounter++

			books = append(books, fmt.Sprintf("Buku %d", currentCounter))

			counter <- currentCounter
			wg.Done()
		}()
	}

	wg.Wait()

	resultCounter := <-counter
	fmt.Printf("Hasil counter sekarang: %d\n", resultCounter)
}
