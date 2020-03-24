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
			err := Render(Setting.TargetFile(name), name)
			if err != nil {
				return err
			}
		}
	}
	// sample
	if Setting.Sample {
		for _, name := range Setting.Samples {
			if strings.HasPrefix(name, Setting.Doc) {
				err := Render(Setting.TargetFile(name), name)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
