package cmd

import (
	"fmt"

	"github.com/Praseetha-KR/strand"
	"github.com/urfave/cli/v2"
)

// URLSafeCmd for generating url-safe string
var URLSafeCmd = &cli.Command{
	Name:    "urlsafe",
	Aliases: []string{"u"},
	Usage:   "Generates a random URL safe string",
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
}
