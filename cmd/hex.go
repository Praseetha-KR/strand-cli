package cmd

import (
	"fmt"

	"github.com/Praseetha-KR/strand"
	"github.com/urfave/cli/v2"
)

// HexCmd for generating hex string
var HexCmd = &cli.Command{
	Name:    "hex",
	Aliases: []string{"h"},
	Usage:   "Generates a random hexadecimal string",
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
}
