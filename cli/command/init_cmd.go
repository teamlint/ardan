package command

import (
	"strings"

	"github.com/urfave/cli/v2"
)

// InitCmd init command layer layout
var InitCmd = &cli.Command{
	Name:   "cmd",
	Usage:  "initial executed command layer layout",
	Action: initCmdCode,
}

func initCmdCode(c *cli.Context) error {
	// code
	for _, name := range Setting.Codes {
		if strings.HasPrefix(name, Setting.Cmd) {
			err := Render(Setting.TargetFile(name), name)
			if err != nil {
				return err
			}
		}
	}
	// sample
	if Setting.Sample {
		for _, name := range Setting.Samples {
			if strings.HasPrefix(name, Setting.Cmd) {
				err := Render(Setting.TargetFile(name), name)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
