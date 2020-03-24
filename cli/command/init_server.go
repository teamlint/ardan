package command

import (
	"strings"

	"github.com/urfave/cli/v2"
)

// InitServer init server layer layout
var InitServer = &cli.Command{
	Name:    "s",
	Aliases: []string{"server"},
	Usage:   "initial server layer layout",
	Action:  initServerCode,
}

func initServerCode(c *cli.Context) error {
	// code
	for _, name := range Setting.Codes {
		if strings.HasPrefix(name, Setting.Server) {
			target := Setting.TargetFile(name)
			err := Render(target, name)
			if err != nil {
				return err
			}
			info(c, "-- %v generated.\n", target)
		}
	}
	// sample
	if Setting.Sample {
		for _, name := range Setting.Samples {
			if strings.HasPrefix(name, Setting.Server) {
				target := Setting.TargetFile(name)
				err := Render(target, name)
				if err != nil {
					return err
				}
				info(c, "-- %v generated.\n", target)
			}
		}
	}

	return nil
}
