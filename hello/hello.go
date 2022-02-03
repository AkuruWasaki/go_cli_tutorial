package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

// Flagのサンプル
// --lang というOptionをつけたら出力される言語が変わる

func main() {
	var language string
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Aliases:     []string{"l"}, // -l を --langの省略形にできる
				Value:       "english",
				Usage:       "language for the greeting",
				EnvVars:     []string{"APP_LANG"},
				Destination: &language,
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					fmt.Println("タスク完了")
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					fmt.Println("タスク追加")
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			name := "BALEEN STUDIO"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if language == "japanese" {
				fmt.Println("こんにちは", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
