package command

import (
	"github.com/teamlint/ardan/cli/setting"
	"github.com/urfave/cli/v2"
)

// GenAll gen all codes
var GenAll = &cli.Command{
	Name:    "all",
	Aliases: []string{"a"},
	Usage:   "generate model all source codes",
	Action:  genAll,
}

func genAll(c *cli.Context) error {
	if err := genModelFile(c, setting.GenTypeAll); err != nil {
		return err
	}

	return nil
}
