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
		return cli.ShowSubcommandHelp(c)
	},
	Subcommands: []*cli.Command{
		InitAll,
		InitLayout,
		InitApp,
		InitCmd,
		InitDoc,
		InitServer,
	},
}
