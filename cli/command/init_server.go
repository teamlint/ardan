package command

import (
	"github.com/urfave/cli/v2"
)

// InitServer init server layer layout
var InitServer = &cli.Command{
	Name:    "server",
	Aliases: []string{"s"},
	Usage:   "initial server layer layout",
	Action:  initServerCode,
}

func initServerCode(c *cli.Context) error {
	// code
	for _, name := range Setting.Codes {
		if Setting.HasPrefix(name, Setting.Server) {
			target := Setting.TargetFile(name)
			err := Render(target, name, nil)
			if err != nil {
				return err
			}
			info(c, "-> %v\n", target)
		}
	}
	// sample
	if Setting.Sample {
		for _, name := range Setting.Samples {
			if Setting.HasPrefix(name, Setting.Server) {
				target := Setting.TargetFile(name)
				err := Render(target, name, nil)
				if err != nil {
					return err
				}
				info(c, "-> %v\n", target)
			}
		}
	}

	return nil
}
