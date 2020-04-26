package cmd

import (
	"fmt"

	"github.com/Praseetha-KR/strand"
	"github.com/urfave/cli/v2"
)

// FromCmd for generating string from charset
var FromCmd = &cli.Command{
	Name:    "from",
	Aliases: []string{"f"},
	Usage:   "Generates a random string from the given character list",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "length",
			Aliases: []string{"l"},
			Value:   10,
			Usage:   "length",
		},
		&cli.StringFlag{
			Name:     "characters",
			Aliases:  []string{"c"},
			Value:    "",
			Required: true,
			Usage:    "character set",
		},
	},
	Action: func(c *cli.Context) error {
		l := validatedLength(c.Int("length"))
		chars := c.String("characters")

		r, err := strand.From(chars, l)
		if err != nil {
			return err
		}
		fmt.Println(r)
		return nil
	},
}
