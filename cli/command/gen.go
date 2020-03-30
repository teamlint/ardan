package command

import (
	"fmt"

	"github.com/teamlint/ardan/cli/setting"
	"github.com/urfave/cli/v2"
)

// Gen sync database tabels struct
var Gen = &cli.Command{
	Name:    "gen",
	Aliases: []string{"generate"},
	Usage:   "generate source codes",
	Action: func(c *cli.Context) error {
		var err error
		err = gen(c)
		if err != nil {
			return err
		}
		return nil
	},
	// Subcommands: []*cli.Command{
	// 	GenController,
	// },
}

func gen(c *cli.Context) error {
	if err := genModelFile(c); err != nil {
		return err
	}
	// run main
	// info(c, ">> %v\n", "generated")

	return nil
}

func genModelFile(c *cli.Context) error {
	models, err := ParseModelFiles(c, setting.DirectiveGen)
	if err != nil {
		return err
	}
	// repository

	// service
	// controller
	for _, name := range Setting.Gens {
		if Setting.IsIteration(name) {
			// iteration
			for _, model := range models {
				// 获取指令参数
				gs, err := Setting.ParseDirectiveGen(model.Directive)
				if err != nil {
					return err
				}
				// info(c, "direc=%v, all=%v, repository=%v, service=%v, controller=%v\n", gs.Directive, gs.All, gs.Repository, gs.Service, gs.Controller)
				if err := genFile(c, name, model, gs); err != nil {
					return err
				}
				// genRepositoryFiel(c, model)
				// target := Setting.TargetFile(name)
				// err = Render(target, name, map[string]interface{}{"Models": beans})
				// if err != nil {
				// 	return fmt.Errorf("sync.Render err=%v\n", err)
				// }
				// // go run
				// output, err := GoRun(filepath.Dir(target), setting.GoMainFile)
				// if err != nil {
				// 	return fmt.Errorf("sync.GoRun err=%v\n", err)
				// }
				// info(c, "--------------------- [sync] start ----------------------------\n"+output)
				// pkg.DeleteFile(target, true)
				// info(c, "---------------------- [sync] end -----------------------------\n")
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
func genFile(c *cli.Context, name string, model *Model, gs *setting.GenSet) error {
	// name := Setting.TmplFile(Setting.App, Setting.Repository, setting.IterationFileTmpl)
	// tfile := filepath.Join(Setting.App, Setting.Repository, strings.ToLower(model.Name)+".go")
	// target := Setting.TargetFile(name)
	target := Setting.TargetFile(name, model.Name)
	info(c, "gen.repository directive=%v, tmpl=%v, target=%v\n", model.Directive, name, target)

	err := Render(target, name, map[string]interface{}{"Model": model})
	if err != nil {
		return fmt.Errorf("ardan.gen err=%v\n", err)
	}
	info(c, ">> %v\n", target)
	return nil
}
