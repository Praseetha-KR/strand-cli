package cmd

import (
	"fmt"

	"github.com/Praseetha-KR/strand"
	"github.com/urfave/cli/v2"
)

// PasswordCmd for generating password string
var PasswordCmd = &cli.Command{
	Name:    "password",
	Aliases: []string{"p"},
	Usage:   "Generates a password of random length",
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
}
