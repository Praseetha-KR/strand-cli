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
				Usage:   "generate random string from alphabet",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "length",
						Aliases: []string{"l"},
						Value:   10,
						Usage:   "length",
					},
					&cli.StringFlag{
						Name:    "case",
						Aliases: []string{"c"},
						Value:   "",
						Usage:   "is lower/upper case letters",
					},
				},
				Action: func(c *cli.Context) error {
					l := validatedLength(c.Int("length"))
					alphacase := validatedCase(c.String("case"))
					r := ""
					if alphacase == LOWER {
						r = strand.AlphaLower(l)
					} else if alphacase == UPPER {
						r = strand.AlphaUpper(l)
					} else {
						r = strand.Alpha(l)
					}
					fmt.Println(r)
					return nil
				},
			},
			{
				Name:    "alphanumeric",
				Aliases: []string{"an"},
				Usage:   "generate random alphanumeric string",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "length",
						Aliases: []string{"l"},
						Value:   10,
						Usage:   "length",
					},
					&cli.StringFlag{
						Name:    "case",
						Aliases: []string{"c"},
						Value:   "",
						Usage:   "is lower/upper case letters",
					},
				},
				Action: func(c *cli.Context) error {
					l := validatedLength(c.Int("length"))
					alphacase := validatedCase(c.String("case"))
					r := ""
					if alphacase == LOWER {
						r = strand.AlphaLowerNumeric(l)
					} else if alphacase == UPPER {
						r = strand.AlphaUpperNumeric(l)
					} else {
						r = strand.AlphaNumeric(l)
					}
					fmt.Println(r)
					return nil
				},
			},
			{
				Name:    "numeric",
				Aliases: []string{"n"},
				Usage:   "generate random number",
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
					r := strand.Numeric(l)
					fmt.Println(r)
					return nil
				},
			},
			{
				Name:    "urlsafe",
				Aliases: []string{"u"},
				Usage:   "generate random URL safe string",
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
					r := strand.URLSafe(l)
					fmt.Println(r)
					return nil
				},
			},
			{
				Name:    "hex",
				Aliases: []string{"h"},
				Usage:   "generate random hexadecimal string",
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
					r := strand.Hex(l)
					fmt.Println(r)
					return nil
				},
			},
			{
				Name:    "binary",
				Aliases: []string{"b"},
				Usage:   "generate random string composed of binary digits",
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
					r := strand.Binary(l)
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
