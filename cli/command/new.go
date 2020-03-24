package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// New application command
var NewCmd = &cli.Command{
	Name:    "new",
	Aliases: []string{"n"},
	Usage:   "new item",
	Action: func(c *cli.Context) error {
		fmt.Println("new command")
		return nil
	},
	Subcommands: []*cli.Command{
		// model
		{
			Name:     "model",
			Usage:    "create application model",
			Category: "app",
			Action: func(c *cli.Context) error {
				fmt.Println("new model command")
				return nil
			},
		},
		// repository
		{
			Name:     "repository",
			Aliases:  []string{"repo"},
			Usage:    "generate application repository",
			Category: "app",
			Action: func(c *cli.Context) error {
				fmt.Println("app repository command")
				return nil

			},
		},
		// service
		{
			Name:     "service",
			Aliases:  []string{"svc"},
			Usage:    "generate application service",
			Category: "app",
			Action: func(c *cli.Context) error {
				fmt.Println("app service command")
				return nil

			},
		},
		// controller
		{
			Name:     "controller",
			Aliases:  []string{"ctrl"},
			Usage:    "generate server controller",
			Category: "server",
			Action: func(c *cli.Context) error {
				fmt.Println("server controller command")
				return nil

			},
		},
		// module
		{
			Name:     "module",
			Aliases:  []string{"mod", "modu"},
			Usage:    "generate server module",
			Category: "server",
			Action: func(c *cli.Context) error {
				fmt.Println("server module command")
				return nil

			},
		},
		// middleware
		{
			Name:     "middleware",
			Aliases:  []string{"mdw", "mid"},
			Usage:    "generate server middleware",
			Category: "server",
			Action: func(c *cli.Context) error {
				fmt.Println("server middleware command")
				return nil

			},
		},
	},
}
