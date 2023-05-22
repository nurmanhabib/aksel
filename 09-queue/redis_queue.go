package queue

import (
	"02-concurrency/09-queue/contract"
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
	client *redis.Client
}

var _ contract.Queue = &RedisQueue{}

func NewRedisQueue(client *redis.Client) contract.Queue {
	return &RedisQueue{client: client}
}

func (r *RedisQueue) Count(ctx context.Context) int64 {
	return r.client.LLen(ctx, "queue").Val()
}

func (r *RedisQueue) Push(ctx context.Context, job *contract.Job) error {
	bytes, err := json.Marshal(job)
	if err != nil {
		return err
	}

	return r.client.LPush(ctx, "queue", bytes).Err()
}

func (r *RedisQueue) Pop(ctx context.Context) (*contract.Job, error) {
	result, err := r.client.RPop(ctx, "queue").Result()
	if err != nil {
		return nil, err
	}

	var job contract.Job

	err = json.Unmarshal([]byte(result), &job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}
