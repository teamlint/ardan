package command

import (
	"github.com/teamlint/ardan/cli/setting"
	"github.com/urfave/cli/v2"
)

// GenRepository gen repository
var GenRepository = &cli.Command{
	Name:    "repository",
	Aliases: []string{"repo"},
	Usage:   "generate repository source codes",
	Action:  genRepo,
}

func genRepo(c *cli.Context) error {
	if err := genModelFile(c, setting.GenTypeRepository); err != nil {
		return err
	}
	return nil
}
