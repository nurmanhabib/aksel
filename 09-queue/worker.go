package queue

import (
	"02-concurrency/09-queue/contract"
	"context"
	"log"
	"time"
)

type Worker struct {
	driver contract.Queue
}

func NewWorker(driver contract.Queue) *Worker {
	return &Worker{driver: driver}
}

func (w *Worker) Run(ctx context.Context) {
	for {
		if w.driver.Count(ctx) > 0 {
			job, err := w.driver.Pop(ctx)
			if err != nil {
				panic(err)
			}

			log.Printf("Ada job masuk dengan: %s, %s", job.ID, job.Message)

			continue
		}

		time.Sleep(5 * time.Second)
	}
}