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
			err := Render(Setting.TargetFile(name), name)
			if err != nil {
				return err
			}
		}
	}
	// sample
	if Setting.Sample {
		for _, name := range Setting.Samples {
			if strings.HasPrefix(name, Setting.Server) {
				err := Render(Setting.TargetFile(name), name)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
