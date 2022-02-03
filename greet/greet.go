package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "greetagain", // コマンド名
		Usage: "fight the loneliness!",
		Action: func(c *cli.Context) error {
			fmt.Println("Good morning friend!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
