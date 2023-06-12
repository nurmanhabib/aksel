package main

import (
	"02-concurrency/10-message-broker/broker"
	"context"

	"github.com/redis/go-redis/v9"
)

func main() {
	var b broker.Broker

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	b = broker.NewRedisDriver(rdb)

	b.Publish(context.Background(), "topic-1", "halo")
	b.Publish(context.Background(), "topic-1", "apa kabar")
	b.Publish(context.Background(), "topic-1", "ini ada pesan di topic 1")
	b.Publish(context.Background(), "topic-2", "ini ada pesan di topic 2 apakah masuk?")
}
