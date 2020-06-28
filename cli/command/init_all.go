package command

import (
	"github.com/urfave/cli/v2"
)

// InitAll initial application project layout & infrastructure codes
var InitAll = &cli.Command{
	Name:    "all",
	Aliases: []string{"a"},
	Usage:   "initial project layout & infrastructure codes",
	Action:  initAll,
}

func initAll(c *cli.Context) error {
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
}
