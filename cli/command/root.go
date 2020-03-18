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
	GoPkgName   string // root package name
	GoModName   string // go module name
	OutputDir   string // output path
	// project layout
	CmdDir          string // cmd directory
	DocDir          string // doc directory
	AppDir          string // application layter directory
	DomainDir       string // domain layter directory
	ServiceDir      string // domain layter directory
	RepositoryDir   string // repository layter directory
	ServerDir       string // server layter directory
	ServerModuleDir string // server module directory
	ServerGlobalDir string // server global directory
	ControllerDir   string // controller directory
	HandlerDir      string // handler directory
	MiddlewareDir   string // middleware directory
	Demo            bool   // demo code
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
				Name:        "conf",
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
				Usage:       "go package name",
				Destination: &GoPkgName,
			},
			&cli.StringFlag{
				Name:        "mod",
				Aliases:     []string{"m"},
				Usage:       "go module name",
				Required:    true,
				Destination: &GoModName,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Value:       "./output",
				Usage:       "output directory",
				Destination: &OutputDir,
			},
			&cli.BoolFlag{
				Name:        "demo",
				Aliases:     []string{"d"},
				Usage:       "initial demo code",
				Destination: &Demo,
			},
			// project layout
			&cli.StringFlag{
				Name:        "cmd",
				Value:       "cmd",
				Usage:       "[Layout] executed command root directory",
				Destination: &CmdDir,
			},
			&cli.StringFlag{
				Name:        "doc",
				Value:       "doc",
				Usage:       "[Layout] documents root directory",
				Destination: &DocDir,
			},
			&cli.StringFlag{
				Name:        "app",
				Value:       "app",
				Usage:       "[Layout] application layer root directory",
				Destination: &AppDir,
			},
			&cli.StringFlag{
				Name:        "domain",
				Value:       "model",
				Usage:       "[Layout] application layer domain directory",
				Destination: &DomainDir,
			},
			&cli.StringFlag{
				Name:        "service",
				Value:       "service",
				Usage:       "[Layout] application layer service directory",
				Destination: &ServiceDir,
			},
			&cli.StringFlag{
				Name:        "repository",
				Value:       "repository",
				Usage:       "[Layout] application layer repository directory",
				Destination: &RepositoryDir,
			},
			&cli.StringFlag{
				Name:        "server",
				Value:       "server",
				Usage:       "[Layout] server layer root directory",
				Destination: &ServerDir,
			},
			&cli.StringFlag{
				Name:        "module",
				Value:       "module",
				Usage:       "[Layout] server layer module directory",
				Destination: &ServerModuleDir,
			},
			&cli.StringFlag{
				Name:        "global",
				Value:       "global",
				Usage:       "[Layout] server layer global directory",
				Destination: &ServerGlobalDir,
			},
			&cli.StringFlag{
				Name:        "controller",
				Value:       "controller",
				Usage:       "[Layout] server layer controller directory",
				Destination: &ControllerDir,
			},
			&cli.StringFlag{
				Name:        "middleware",
				Value:       "middleware",
				Usage:       "[Layout] server layer middleware directory",
				Destination: &MiddlewareDir,
			},
			&cli.StringFlag{
				Name:        "handler",
				Value:       "handler",
				Usage:       "[Layout] server layer handler directory",
				Destination: &HandlerDir,
			},
		},
		Action: func(c *cli.Context) error {
			cli.ShowAppHelp(c)
			// cli.ShowSubcommandHelp(c)
			fmt.Printf("templates=%v\n", TemplateDir)
			fmt.Printf("config=%v\n", ConfigFile)
			fmt.Printf("conn=%v\n", ConnStr)
			fmt.Printf("goPkgName=%v\n", GoPkgName)
			fmt.Printf("goModName=%v\n", GoModName)
			fmt.Printf("output=%v\n", OutputDir)
			fmt.Printf("demo=%v\n", Demo)
			return nil
		},
		Commands: []*cli.Command{
			App,
			Init,
		},
	}

	return app.Run(os.Args)
}
