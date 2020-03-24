package command

import (
	"github.com/urfave/cli/v2"
)

// Init init project layout
var Init = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "initial project layout",
	Action: func(c *cli.Context) error {
		var err error
		// layout
		err = initLayout(c)
		if err != nil {
			return err
		}
		// app
		err = initAppCode(c)
		if err != nil {
			return err
		}
		// cmd
		err = initCmdCode(c)
		if err != nil {
			return err
		}
		// doc
		err = initDocCode(c)
		if err != nil {
			return err
		}
		// server
		err = initServerCode(c)
		if err != nil {
			return err
		}
		return nil
	},
	Subcommands: []*cli.Command{
		InitLayout,
		InitApp,
		InitCmd,
		InitDoc,
		InitServer,
	},
}
