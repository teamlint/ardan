package command

import (
	"fmt"

	"github.com/rakyll/statik/fs"
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
		b, err := fs.ReadFile(Setting.FileSystem, o)
		if err != nil {
			return fmt.Errorf("res.ReadFile file=%v err=%v\n", o, err)
		}
		// log.Printf("res.ReadFile src=%v content=%v\n", o, string(b))
		dst := Setting.TargetFile(o)
		if err := pkg.WriteFile(dst, b); err != nil {
			return fmt.Errorf("res.WriteFile src=%v dest=%v err=%v\n", o, dst, err)
		}
		info(c, "=> %v\n", dst)
	}
	return nil
}
