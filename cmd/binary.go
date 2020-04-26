package cmd

import (
	"fmt"

	"github.com/Praseetha-KR/strand"
	"github.com/urfave/cli/v2"
)

// BinaryCmd for generating binary string
var BinaryCmd = &cli.Command{
	Name:    "binary",
	Aliases: []string{"b"},
	Usage:   "Generates a random string of binary digits",
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
}
