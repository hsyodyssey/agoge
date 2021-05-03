package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hsyodyssey/agoge/router"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()

	// Set Flags
	// Test    env: --env test[default]
	// Product env: --env product
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "env",
			Value:    "test",
			Usage:    "Start the Gin server",
			Required: true,
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.String("env") == "product" {
			fmt.Println("[AGOGE] The product env is Staring...")
			fmt.Println("[AGOGE] Listening on 8080...")
			r := router.SetupRouter()
			r.Run(":8080")
		} else if c.String("env") == "test" {
			fmt.Println("[AGOGE-TEST] The Test env is Staring...")
			fmt.Println("[AGOGE-TEST] Listening on 8000...")
			r := router.SetupRouter()
			r.Run(":8000")
		} else {
			fmt.Println("Please input the right flags [like: test or product]")
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
