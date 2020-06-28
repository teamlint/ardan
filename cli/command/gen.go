package command

import (
	"fmt"

	"github.com/teamlint/ardan/cli/setting"
	"github.com/urfave/cli/v2"
)

// Gen sync database tabels struct
var Gen = &cli.Command{
	Name:    "generate",
	Aliases: []string{"gen", "g"},
	Usage:   "generate source codes",
	Action: func(c *cli.Context) error {
		return cli.ShowSubcommandHelp(c)
	},
	Subcommands: []*cli.Command{
		GenAll,
		GenQuery,
		GenRepository,
		GenService,
		GenController,
	},
}

func genModelFile(c *cli.Context, genType setting.GenType) error {
	models, err := ParseModelFiles(c, setting.DirectiveGen)
	if err != nil {
		return err
	}
	for _, name := range Setting.Gens {
		if Setting.IsIteration(name) {
			// iteration
			for _, model := range models {
				if err := genFile(c, name, model, genType); err != nil {
					return err
				}
			}
		} else {
			// single
			target := Setting.TargetFile(name)
			//
			err := Render(target, name, map[string]interface{}{"Models": models})
			if err != nil {
				return err
			}
			info(c, ">> %v\n", target)
		}
	}

	return nil
}
func genFile(c *cli.Context, name string, model *Model, genType setting.GenType) error {
	var target string
	// info(c, "gen.%s tmpl=%s\n", genType, name)
	switch genType {
	case setting.GenTypeQuery: // query
		if !Setting.IsQuery(name) {
			return nil
		}
	case setting.GenTypeRepository: // repository
		if !Setting.IsRepository(name) {
			return nil
		}
	case setting.GenTypeService: // service

		if !Setting.IsService(name) {
			return nil
		}
	case setting.GenTypeController: // controller

		if !Setting.IsController(name) {
			return nil
		}
	case setting.GenTypeAll: // all
	default:
		info(c, "not found <ardan:gen> directive\n")
		return nil
	}
	target = Setting.TargetFile(name, setting.SnakeCase(model.Name))
	// info(c, "gen %s tmpl=%s, target=%s\n", genType, name, target)
	err := Render(target, name, map[string]interface{}{"Model": model})
	if err != nil {
		return fmt.Errorf("gen %s %v err=%v", genType, target, err)
	}
	info(c, ">> %v\n", target)
	return nil
}
