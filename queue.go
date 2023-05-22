package main

import (
	queue "02-concurrency/09-queue"
	"02-concurrency/09-queue/contract"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/aksel?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	driver := queue.NewDBQueue(db)

	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(*cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "push",
				Usage: "add a task to the list",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					driver.Push(cCtx.Context, &contract.Job{
						ID:      uuid.New().String(),
						Message: cCtx.Args().First(),
					})

					return nil
				},
			},
			{
				Name:  "work",
				Usage: "Run a worker",
				Action: func(cCtx *cli.Context) error {
					worker := queue.NewWorker(driver)

					worker.Run(context.Background())

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
