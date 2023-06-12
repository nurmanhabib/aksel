package main

import (
	"02-concurrency/10-message-broker/broker"
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

func main() {
	var b broker.Broker

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	b = broker.NewRedisDriver(rdb)

	var wg sync.WaitGroup

	ctx := context.Background()

	wg.Add(1)
	go func() {
		defer wg.Done()
		b.Subscribe(ctx, "topic-1", func(data interface{}) {
			log.Printf("Ada pesan masuk dari topic-1: %v", data)
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		b.Subscribe(ctx, "topic-2", func(data interface{}) {
			log.Printf("Ada pesan masuk dari topic-2: %v", data)
		})
	}()

	wg.Wait()

	log.Println("Program berakhir...")
}
