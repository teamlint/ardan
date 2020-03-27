package command

import (
	"path/filepath"
	"strings"

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
	info(c, ">> %v\n", "generated")

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
	for _, model := range models {
		genRepositoryFiel(c, model)
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

	return nil
}
func genRepositoryFiel(c *cli.Context, model *Model) error {
	name := Setting.TmplFile(Setting.App, Setting.Repository, setting.GoGenFileTmpl)
	tfile := filepath.Join(Setting.App, Setting.Repository, strings.ToLower(model.Name)+".go")
	target := Setting.TargetFile(tfile)
	info(c, "gen.repository directive=%v, tmpl=%v, target=%v\n", model.Directive, name, target)
	return nil
}
