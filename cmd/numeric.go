package cmd

import (
	"fmt"

	"github.com/Praseetha-KR/strand"
	"github.com/urfave/cli/v2"
)

// NumericCmd for generating numeric string
var NumericCmd = &cli.Command{
	Name:    "numeric",
	Aliases: []string{"n"},
	Usage:   "Generates a random number",
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
}
