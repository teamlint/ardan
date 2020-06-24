package command

import (
	"github.com/teamlint/ardan/cli/setting"
	"github.com/urfave/cli/v2"
)

// GenController gen controller
var GenController = &cli.Command{
	Name:    "controller",
	Aliases: []string{"ctrl"},
	Usage:   "generate controller source codes",
	Action:  genController,
}

func genController(c *cli.Context) error {
	if err := genModelFile(c, setting.GenTypeController); err != nil {
		return err
	}
	return nil
}
