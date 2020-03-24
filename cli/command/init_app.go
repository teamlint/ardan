package command

import (
	"strings"

	"github.com/urfave/cli/v2"
)

// InitApp init application layer layout
var InitApp = &cli.Command{
	Name:    "app",
	Aliases: []string{"a"},
	Usage:   "initial application layer layout",
	Action:  initAppCode,
}

func initAppCode(c *cli.Context) error {
	// code
	for _, name := range Setting.Codes {
		// log.Printf("[init.app] path=%v,appDir=%v\n", name, Setting.App)
		if strings.HasPrefix(name, Setting.App) {
			// tmpl
			target := Setting.TargetFile(name)
			err := Render(target, name)
			if err != nil {
				return err
			}
			info(c, "-> %v\n", target)
		}
	}
	// sample
	if Setting.Sample {
		for _, name := range Setting.Samples {
			if strings.HasPrefix(name, Setting.App) {
				// tmpl
				target := Setting.TargetFile(name)
				err := Render(target, name)
				if err != nil {
					return err
				}
				info(c, "-> %v\n", target)
			}
		}
	}

	return nil
}
