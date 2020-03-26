package command

import (
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
		if Setting.HasPrefix(name, Setting.Doc) {
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
			if Setting.HasPrefix(name, Setting.Doc) {
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
