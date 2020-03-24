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
	// layout directory
	for _, l := range Setting.Layouts {
		if err := pkg.Mkdir(l); err != nil {
			return err
		}
		info(c, "*layout* = %v\n", l)
	}
	// origin files
	for _, o := range Setting.Origins {
		dst := Setting.TargetFile(o)
		if err := pkg.Copy(Setting.SourceFile(o), dst); err != nil {
			return err
		}
		info(c, "=> %v\n", dst)
	}
	return nil
}
