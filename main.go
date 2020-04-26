package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Praseetha-KR/strand"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "strand",
		Usage: "Random String Generator",
		Commands: []*cli.Command{
			{
				Name:    "string",
				Aliases: []string{"s"},
				Usage:   "generate random string",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "length",
						Aliases: []string{"l"},
						Value:   10,
						Usage:   "length",
					},
				},
				Action: func(c *cli.Context) error {
					l := validatedLength(c.Int("length"))
					r := strand.String(l)
					fmt.Println(r)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
