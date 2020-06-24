package command

import (
	"github.com/teamlint/ardan/cli/setting"
	"github.com/urfave/cli/v2"
)

// GenQuery gen model query
var GenQuery = &cli.Command{
	Name:    "query",
	Aliases: []string{"q"},
	Usage:   "generate model query source codes",
	Action:  genQuery,
}

func genQuery(c *cli.Context) error {
	if err := genModelFile(c, setting.GenTypeQuery); err != nil {
		return err
	}
	return nil
}
