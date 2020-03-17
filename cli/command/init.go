package command

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/teamlint/ardan/cli/lib"
	"github.com/teamlint/ardan/cli/setting"
	"github.com/urfave/cli/v2"
)

// Init init project layout
var Init = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "initial project layout",
	Action: func(c *cli.Context) error {
		fmt.Println("init command")
		fmt.Printf("templates=%v\n", TemplateDir)
		fmt.Printf("pkgName=%v\n", PkgName)
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
	Action: func(c *cli.Context) error {
		// init setting
		setting.Init(TemplateDir, OutputDir)
		// 遍历模板目录,生成目标目录结构
		err := filepath.Walk(TemplateDir, walkTmplDir)
		log.Printf("current pkgname=%v\n", lib.GetPkgName())
		set := setting.Instance()
		log.Printf("template engine=%+v appDir=%v\n", *set.Template, set.AppDir)

		return err
	},
}

func walkTmplDir(path string, info os.FileInfo, err error) error {
	fmt.Printf("path=%v,is_dir=%v\n", path, info.IsDir())
	if info.IsDir() {
		err := os.MkdirAll(filepath.Join(OutputDir, path), os.ModePerm)
		return err
	}
	return nil
}
