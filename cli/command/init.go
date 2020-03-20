package command

import (
	"log"
	"strings"

	"github.com/teamlint/ardan/cli/lib"
	"github.com/urfave/cli/v2"
)

// Init init project layout
var Init = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "initial project layout",
	Action: func(c *cli.Context) error {
		var err error
		// layout
		err = initLayout(c)
		if err != nil {
			return err
		}
		// codes
		err = initAppCode(c)
		if err != nil {
			return err
		}
		// sample
		return nil
	},
	Subcommands: []*cli.Command{
		InitLayout,
		InitApp,
	},
}

// InitLayout initial application project layout
var InitLayout = &cli.Command{
	Name:    "layout",
	Aliases: []string{"l"},
	Usage:   "initial project layout",
	Action:  initLayout,
}

func initLayout(c *cli.Context) error {
	for _, l := range Setting.Layouts {
		log.Printf("[init.layout]=%v\n", l)
		if err := lib.Mkdir(l); err != nil {
			return err
		}
	}
	return nil
}

// InitApp init application layer layout
var InitApp = &cli.Command{
	Name:    "app",
	Aliases: []string{"a"},
	Usage:   "initial application layer layout",
	Action:  initAppCode,
}

func initAppCode(c *cli.Context) error {
	t := Setting.Template
	// code
	for _, name := range Setting.Codes {
		log.Printf("[init.app] path=%v,appDir=%v\n", name, AppDir)
		if strings.HasPrefix(name, AppDir) {
			f, err := lib.NewFile(Setting.TargetFile(name))
			if err != nil {
				log.Printf("error creating generated code %v: %v\n", name, err)
				return err
			}
			defer f.Close()
			// tmpl
			err = t.ExecuteTemplate(f, name, tmplData())
			if err != nil {
				log.Printf("template render code err=%v\n", err)
				return err
			}
		}
	}
	// sample
	if Sample {
		for _, name := range Setting.Samples {
			if strings.HasPrefix(name, AppDir) {
				f, err := lib.NewFile(Setting.TargetFile(name))
				if err != nil {
					log.Printf("error creating generated sample %v: %v\n", name, err)
					return err
				}
				defer f.Close()
				// tmpl
				err = t.ExecuteTemplate(f, name, tmplData())
				if err != nil {
					log.Printf("template render sample err=%v\n", err)
					return err
				}
			}
		}
	}

	return nil
}

func tmplData() map[string]interface{} {
	data := make(map[string]interface{})
	data["Setting"] = Setting
	return data
}
