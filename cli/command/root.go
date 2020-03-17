package command

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	TemplateDir string // template file root path
	ConfigFile  string // config file path
	ConnStr     string // database connection string
	PkgName     string // root package name
	ModuleName  string // go module name
	OutputDir   string // output path
)

// Run command run
func Run() error {
	app := &cli.App{
		Name:    "ardan",
		Usage:   "make an app bootstrap",
		Version: "v1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "tmpl",
				Aliases:     []string{"t"},
				Value:       "templates",
				Usage:       "template root dir",
				Destination: &TemplateDir,
			},
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"f"},
				Value:       "",
				Usage:       "config file path",
				Destination: &ConfigFile,
			},
			&cli.StringFlag{
				Name:        "conn",
				Aliases:     []string{"c"},
				Usage:       "database connection string",
				Destination: &ConnStr,
			},
			&cli.StringFlag{
				Name:        "pkg",
				Aliases:     []string{"p"},
				Usage:       "package name",
				Destination: &PkgName,
			},
			&cli.StringFlag{
				Name:        "module",
				Aliases:     []string{"m"},
				Usage:       "go module name",
				Destination: &ModuleName,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Value:       "./output",
				Usage:       "output dir",
				Destination: &OutputDir,
			},
		},
		Action: func(c *cli.Context) error {
			cli.ShowAppHelp(c)
			fmt.Printf("templates=%v\n", TemplateDir)
			fmt.Printf("config=%v\n", ConfigFile)
			fmt.Printf("conn=%v\n", ConnStr)
			fmt.Printf("pkgName=%v\n", PkgName)
			fmt.Printf("output=%v\n", OutputDir)
			return nil
		},
		Commands: []*cli.Command{
			App,
			Init,
		},
	}

	return app.Run(os.Args)
}
