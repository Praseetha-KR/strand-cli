package cmd

import (
	"fmt"

	"github.com/Praseetha-KR/strand"
	"github.com/urfave/cli/v2"
)

// AlphaNumericCmd for generating alphanumeric string
var AlphaNumericCmd = &cli.Command{
	Name:    "alphanumeric",
	Aliases: []string{"an"},
	Usage:   "Generates a random alphanumeric string",
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
}
