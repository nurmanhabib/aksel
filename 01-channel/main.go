package main

import "fmt"

func main() {
	// Membuat channel (non-buffered channel) / tidak punya kapasitas channel
	notification := make(chan string)

	// Goroutine - mengirim pesan via channel
	go func() {
		notification <- "hai"
	}()

	go func() {
		notification <- "apa kabar"
	}()

	fmt.Printf("ada pesan: %s\n", <-notification)
	fmt.Printf("ada pesan: %s\n", <-notification)
}
