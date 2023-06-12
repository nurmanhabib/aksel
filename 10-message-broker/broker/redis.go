package broker

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisDriver struct {
	client *redis.Client
}

// Type Assertion
var _ Broker = &RedisDriver{}

func NewRedisDriver(client *redis.Client) Broker {
	return &RedisDriver{client}
}

func (r *RedisDriver) Publish(ctx context.Context, topic string, data interface{}) error {
	// result := r.client.Publish(ctx, topic, data)
	// return result.Err()

	// result := r.client.Publish(ctx, topic, data)
	//
	// if result.Err() != nil {
	// 	return result.Err()
	// }

	// return nil

	return r.client.Publish(ctx, topic, data).Err()
}

func (r *RedisDriver) Subscribe(ctx context.Context, topic string, receiver Receiver) {
	subs := r.client.Subscribe(ctx, topic)

	for {
		message, err := subs.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		receiver(message)
	}
}
