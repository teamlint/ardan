package command

import (
	"github.com/teamlint/ardan/pkg"
	"github.com/urfave/cli/v2"
)

// InitLayout initial application project layout
var InitLayout = &cli.Command{
	Name:    "layout",
	Aliases: []string{"l"},
	Usage:   "initial project layout",
	Action:  initLayout,
}

func initLayout(c *cli.Context) error {
	for _, l := range Setting.Layouts {
		info(c, "[layout] = %v\n", l)
		if err := pkg.Mkdir(l); err != nil {
			return err
		}
	}
	return nil
}
