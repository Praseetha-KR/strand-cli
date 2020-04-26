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
			{
				Name:    "alpha",
				Aliases: []string{"a"},
				Usage:   "generate random number",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "length",
						Aliases: []string{"l"},
						Value:   10,
						Usage:   "length",
					},
					&cli.StringFlag{
						Name:  "case",
						Value: "",
						Usage: "is lower/upper case letters",
					},
				},
				Action: func(c *cli.Context) error {
					l := validatedLength(c.Int("length"))
					alphacase := validatedCase(c.String("case"))
					r := ""
					if alphacase == "LOWER" {
						r = strand.AlphaLower(l)
					} else if alphacase == "UPPER" {
						r = strand.AlphaUpper(l)
					} else {
						r = strand.Alpha(l)
					}
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
