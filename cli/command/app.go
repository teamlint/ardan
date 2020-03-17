package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// App application command
var App = &cli.Command{
	Name:    "app",
	Aliases: []string{"a"},
	Usage:   "application layer",
	Action: func(c *cli.Context) error {
		fmt.Println("app command")
		return nil
	},
	Subcommands: []*cli.Command{
		// model
		{
			Name:     "model",
			Usage:    "sync model to database",
			Category: "database",
			Action: func(c *cli.Context) error {
				fmt.Println("app model command")
				return nil
			},
		},
		// repository
		{
			Name:  "repo",
			Usage: "generate application repository",
			Action: func(c *cli.Context) error {
				fmt.Println("app repository command")
				return nil

			},
		},
		// service
	},
}
