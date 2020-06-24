package command

import (
	"github.com/teamlint/ardan/cli/setting"
	"github.com/urfave/cli/v2"
)

// GenService gen service
var GenService = &cli.Command{
	Name:    "service",
	Aliases: []string{"svc"},
	Usage:   "generate service source codes",
	Action:  genService,
}

func genService(c *cli.Context) error {
	if err := genModelFile(c, setting.GenTypeService); err != nil {
		return err
	}
	return nil
}
