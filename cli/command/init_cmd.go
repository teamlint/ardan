package command

import (
	"github.com/urfave/cli/v2"
)

// InitCmd init command layer infrastructure codes
var InitCmd = &cli.Command{
	Name:   "cmd",
	Usage:  "initial executed command layer infrastructure codes",
	Action: initCmdCode,
}

func initCmdCode(c *cli.Context) error {
	// code
	for _, name := range Setting.Codes {
		if Setting.HasPrefix(name, Setting.Cmd) {
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
			if Setting.HasPrefix(name, Setting.Cmd) {
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
