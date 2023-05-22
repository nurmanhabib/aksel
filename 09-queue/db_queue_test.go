package queue_test

import (
	queue "02-concurrency/09-queue"
	"02-concurrency/09-queue/contract"
	"context"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestDBQueue(t *testing.T) {
	dsn := "root:@tcp(127.0.0.1:3306)/aksel?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&queue.Antrian{})

	ctx := context.Background()

	dbQueue := queue.NewDBQueue(db)
	dbQueue.Push(ctx, &contract.Job{
		ID:      "1",
		Message: "hello",
	})

	dbQueue.Push(ctx, &contract.Job{
		ID:      "2",
		Message: "world",
	})

	// job, err := dbQueue.Pop(ctx)
	// require.NoError(t, err)
	//
	// assert.Equal(t, job.ID, "1")
	// assert.Equal(t, job.Message, "hello")
	//
	// job, err = dbQueue.Pop(ctx)
	// require.NoError(t, err)
	//
	// assert.Equal(t, job.ID, "2")
	// assert.Equal(t, job.Message, "world")
}
