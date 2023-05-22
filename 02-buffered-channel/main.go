package main

import "fmt"

func main() {
	fmt.Println("Belajar Goroutine")
	fmt.Printf("Pertemuan Minggu ke-%d Sesi ke-%s\n", 1, "dua")

	// Membuat buffered channel berkapasitas 5
	notification := make(chan string, 2)

	// Send message to notification channel
	notification <- "message 1"
	notification <- "message 2"

	// Menerima message dari notification channel
	fmt.Printf("telah diterima pesan: %s\n", <-notification)
	fmt.Printf("telah diterima pesan: %s\n", <-notification)

	notification <- "message 3"
	fmt.Printf("telah diterima pesan: %s\n", <-notification)
}
