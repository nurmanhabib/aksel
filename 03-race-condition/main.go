package main

import (
	"fmt"
	"sync"
)

func main() {
	// Inisiasi variable counter 0
	counter := 0

	// WaitGroup
	var wg sync.WaitGroup

	// =======================
	// Mendaftarkan sebanyak 1 goroutine
	wg.Add(1)

	// Increment counter dengan goroutine
	go func() {
		// Proses increment
		counter++

		// Ngasih informasi kw waitgroup bahwa proses selesai
		wg.Done()
	}()

	// ========================

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}

	// Menunggu sampai goroutine selesai
	// time.Sleep(1 * time.Second)
	wg.Wait()

	fmt.Printf("Counter sekarang: %d\n", counter)
}
