package command

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	TemplateDir string // template file root path
	ConfigFile  string // config file path
	DBDriver    string // database driver name
	DBName      string // database name
	DBConnStr   string // database connection string
	GoCmd       string // go command name
	GoOS        string // sets the GOOS when producing a binary with -compileout
	GoARCH      string // sets the GOARCH when producing a binary with -compileout
	GoPkgName   string // root package name
	GoModName   string // go module name
	OutputDir   string // output path
	// project layout
	CmdDir          string // cmd directory
	DocDir          string // doc directory
	AppDir          string // application layter directory
	ModelDir        string // domain layter directory
	ServiceDir      string // domain layter directory
	RepositoryDir   string // repository layter directory
	ServerDir       string // server layter directory
	ServerModuleDir string // server module directory
	ServerGlobalDir string // server global directory
	ControllerDir   string // controller directory
	HandlerDir      string // handler directory
	MiddlewareDir   string // middleware directory
	Sample          bool   // sample code
)

// Run command run
func Run() error {
	app := &cli.App{
		Name:         "ardan",
		Usage:        "make an app bootstrap",
		Version:      "v1.0.0",
		Flags:        flags(),
		BashComplete: bashComplete,
		Before:       before,
		After:        after,
		Action: func(c *cli.Context) error {
			cli.ShowAppHelp(c)
			// cli.ShowSubcommandHelp(c)
			info()
			return nil
		},
		Commands: []*cli.Command{
			App,
			Init,
			Sync,
		},
	}

	return app.Run(os.Args)
}

func flags() []cli.Flag {
	return []cli.Flag{
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
			Name:        "db-driver",
			Aliases:     []string{"dd"},
			Value:       "postgres",
			Usage:       "database driver name: mysql|postgres",
			Destination: &DBDriver,
		},
		&cli.StringFlag{
			Name:        "db-name",
			Aliases:     []string{"dn"},
			Value:       "Ardan",
			Usage:       "database name",
			Destination: &DBName,
		},
		&cli.StringFlag{
			Name:        "db-conn",
			Aliases:     []string{"dc"},
			Usage:       "database connection string",
			Destination: &DBConnStr,
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
			Value:       ".",
			Usage:       "go module name",
			Required:    false,
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
			Name:        "sample",
			Aliases:     []string{"s"},
			Value:       false,
			Usage:       "initial sample code",
			Destination: &Sample,
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
			Name:        "model",
			Value:       "model",
			Usage:       "[Layout] application layer domain directory",
			Destination: &ModelDir,
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
	}
}

func bashComplete(c *cli.Context) {
	fmt.Fprintf(c.App.Writer, "bashComplete\n")
}

func before(c *cli.Context) error {
	// fmt.Fprintf(c.App.Writer, "before\n")
	var err error
	err = Setup(c)
	return err
}

func after(c *cli.Context) error {
	fmt.Fprintf(c.App.Writer, "after\n")
	return nil
}
