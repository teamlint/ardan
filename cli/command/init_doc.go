package command

import (
	"strings"

	"github.com/urfave/cli/v2"
)

// InitDoc init document directory
var InitDoc = &cli.Command{
	Name:   "cmd",
	Usage:  "initial document directory",
	Action: initDocCode,
}

func initDocCode(c *cli.Context) error {
	// code
	for _, name := range Setting.Codes {
		if strings.HasPrefix(name, Setting.Doc) {
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
			if strings.HasPrefix(name, Setting.Doc) {
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
