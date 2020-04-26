package cmd

import (
	"fmt"

	"github.com/Praseetha-KR/strand"
	"github.com/urfave/cli/v2"
)

// StringCmd for generating random string
var StringCmd = &cli.Command{
	Name:    "string",
	Aliases: []string{"s"},
	Usage:   "Generates a random string",
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
}
