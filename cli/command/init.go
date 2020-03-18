package command

import (
	"fmt"
	"log"

	"github.com/teamlint/ardan/cli/setting"
	"github.com/urfave/cli/v2"
)

// Init init project layout
var Init = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "initial project layout",
	Action: func(c *cli.Context) error {
		fmt.Println("init root command")
		var err error
		err = initApp(c)
		if err != nil {
			return err
		}
		return nil
	},
	Subcommands: []*cli.Command{
		InitApp,
	},
}

// InitApp init application layer layout
var InitApp = &cli.Command{
	Name:    "app",
	Aliases: []string{"a"},
	Usage:   "initial application layer layout",
	Action:  initApp,
}

func initApp(c *cli.Context) error {
	// init setting
	setting.Init(setting.Options{
		TmplDir:   TemplateDir,
		OutputDir: OutputDir,
		Demo:      Demo,
	})
	set := setting.Instance()
	log.Printf("template engine=%+v appDir=%v\n", *set.Template, set.AppDir)

	return nil
}
