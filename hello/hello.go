package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// Flagのサンプル
// --lang というOptionをつけたら出力される言語が変わる

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
		},
		Action: func(c *cli.Context) error {
			name := "BALEEN STUDIO"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.String("lang") == "japanese" {
				fmt.Println("こんにちは", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
