package cmd

import (
	"fmt"

	"github.com/Praseetha-KR/strand"
	"github.com/urfave/cli/v2"
)

// AlphaCmd for generating alpha string
var AlphaCmd = &cli.Command{
	Name:    "alpha",
	Aliases: []string{"a"},
	Usage:   "Generates a random string of letters",
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
}
