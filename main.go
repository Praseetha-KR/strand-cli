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
					&cli.BoolFlag{
						Name:    "lower",
						Aliases: []string{"low"},
						Value:   false,
						Usage:   "is lower",
					},
					&cli.BoolFlag{
						Name:    "upper",
						Aliases: []string{"up"},
						Value:   false,
						Usage:   "is upper",
					},
				},
				Action: func(c *cli.Context) error {
					l := validatedLength(c.Int("length"))
					isLower := c.Bool("lower")
					isUpper := c.Bool("upper")
					r := ""
					if isLower && !isUpper {
						r = strand.AlphaLower(l)
					} else if isUpper && !isLower {
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
					&cli.BoolFlag{
						Name:    "lower",
						Aliases: []string{"low"},
						Value:   false,
						Usage:   "is lower",
					},
					&cli.BoolFlag{
						Name:    "upper",
						Aliases: []string{"up"},
						Value:   false,
						Usage:   "is upper",
					},
				},
				Action: func(c *cli.Context) error {
					l := validatedLength(c.Int("length"))
					isLower := c.Bool("lower")
					isUpper := c.Bool("upper")
					r := ""
					if isLower && !isUpper {
						r = strand.AlphaLowerNumeric(l)
					} else if isUpper && !isLower {
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
			{
				Name:    "password",
				Aliases: []string{"p"},
				Usage:   "generate password of random length",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "simple",
						Aliases: []string{"s"},
						Value:   false,
						Usage:   "8 character password",
					},
					&cli.BoolFlag{
						Name:    "urlsafe",
						Aliases: []string{"u"},
						Value:   false,
						Usage:   "is urlsafe",
					},
				},
				Action: func(c *cli.Context) error {
					isSimple := c.Bool("simple")
					isURLSafe := c.Bool("urlsafe")
					r := ""
					if isURLSafe && isSimple {
						r = strand.URLSafeSimplePassword()
					} else if isURLSafe && !isSimple {
						r = strand.URLSafePassword()
					} else if !isURLSafe && isSimple {
						r = strand.SimplePassword()
					} else {
						r = strand.Password()
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
