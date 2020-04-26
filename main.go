package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Praseetha-KR/strand-cli/cmd"
	"github.com/urfave/cli/v2"
)

const (
	appName     = "strand"
	description = "Random String Generator"
	version     = "v0.0.3"
	author      = "Praseetha KR"
	email       = "praseetha04@gmail.com"
	header      = `
            $$\                                        $$\
            $$ |                                       $$ |
 $$$$$$$\ $$$$$$\    $$$$$$\  $$$$$$\  $$$$$$$\   $$$$$$$ |
$$  _____|\_$$  _|  $$  __$$\ \____$$\ $$  __$$\ $$  __$$ |
\$$$$$$\    $$ |    $$ |  \__|$$$$$$$ |$$ |  $$ |$$ /  $$ |
 \____$$\   $$ |$$\ $$ |     $$  __$$ |$$ |  $$ |$$ |  $$ |
$$$$$$$  |  \$$$$  |$$ |     \$$$$$$$ |$$ |  $$ |\$$$$$$$ |
\_______/    \____/ \__|      \_______|\__|  \__| \_______|


`
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
	cli.AppHelpTemplate = fmt.Sprintf(`%s%s`, header, cli.AppHelpTemplate)
	return app
}
