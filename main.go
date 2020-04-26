package main

import (
	"log"
	"os"

	"github.com/Praseetha-KR/strand-cli/cmd"
	"github.com/urfave/cli/v2"
)

const (
	appName     = "strand"
	description = "Random String Generator"
	version     = "0.0.1"
	author      = "Praseetha KR"
	email       = "praseetha04@gmail.com"
)

func main() {
	if err := newApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newApp() *cli.App {
	app := &cli.App{
		Name:    appName,
		Usage:   description,
		Version: version,
		Authors: []*cli.Author{
			&cli.Author{
				Name:  author,
				Email: email,
			},
		},
		Commands: []*cli.Command{
			cmd.StringCmd,
			cmd.AlphaCmd,
			cmd.AlphaNumericCmd,
			cmd.NumericCmd,
			cmd.URLSafeCmd,
			cmd.HexCmd,
			cmd.BinaryCmd,
			cmd.PasswordCmd,
			cmd.FromCmd,
		},
	}
	return app
}
