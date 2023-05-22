package queue_test

import (
	queue "02-concurrency/09-queue"
	"02-concurrency/09-queue/contract"
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRedisQueue(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   1,
	})

	ctx := context.Background()

	redisQueue := queue.NewRedisQueue(client)
	redisQueue.Push(ctx, &contract.Job{
		ID:      "1",
		Message: "hello",
	})

	redisQueue.Push(ctx, &contract.Job{
		ID:      "2",
		Message: "world",
	})

	job, err := redisQueue.Pop(ctx)
	require.NoError(t, err)

	assert.Equal(t, job.ID, "1")
	assert.Equal(t, job.Message, "hello")

	job, err = redisQueue.Pop(ctx)
	require.NoError(t, err)

	assert.Equal(t, job.ID, "2")
	assert.Equal(t, job.Message, "world")
}
