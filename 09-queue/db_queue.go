package queue

import (
	"02-concurrency/09-queue/contract"
	"context"
	"encoding/json"

	"gorm.io/gorm"
)

type Antrian struct {
	ID      int    `gorm:"primaryKey;autoIncrement"`
	Message string `gorm:"message"`
}

type DBQueue struct {
	db *gorm.DB
}

func NewDBQueue(db *gorm.DB) contract.Queue {
	return &DBQueue{db: db}
}

func (d *DBQueue) Count(ctx context.Context) int64 {
	var count int64

	err := d.db.WithContext(ctx).Table("antrians").Count(&count).Error
	if err != nil {
		return 0
	}

	return count
}

func (d *DBQueue) Push(ctx context.Context, job *contract.Job) error {
	bytes, err := json.Marshal(job)
	if err != nil {
		return err
	}

	antrian := &Antrian{
		Message: string(bytes),
	}

	return d.db.WithContext(ctx).Create(antrian).Error
}

func (d *DBQueue) Pop(ctx context.Context) (*contract.Job, error) {
	var antrian Antrian

	err := d.db.WithContext(ctx).Order("id ASC").First(&antrian).Error
	if err != nil {
		return nil, err
	}

	var job contract.Job

	err = json.Unmarshal([]byte(antrian.Message), &job)
	if err != nil {
		return nil, err
	}

	err = d.db.WithContext(ctx).Delete(&antrian).Error
	if err != nil {
		return nil, err
	}

	return &job, nil
}
