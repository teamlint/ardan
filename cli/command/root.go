package command

import (
	"os"

	"github.com/teamlint/ardan/cli/setting"
	"github.com/urfave/cli/v2"
)

var opts setting.Options

// Run command run
func Run() error {
	app := &cli.App{
		Name:         "ardan",
		Authors:      []*cli.Author{&cli.Author{Name: "venjiang", Email: "venjiang@gmail.com"}},
		Copyright:    "Copyright teamlint.com",
		Usage:        "make an app bootstrap",
		Version:      "v1.0.0",
		Flags:        flags(),
		BashComplete: bashComplete,
		Before:       before,
		After:        after,
		Action: func(c *cli.Context) error {
			cli.ShowAppHelp(c)
			// cli.ShowSubcommandHelp(c)
			return nil
		},
		Commands: []*cli.Command{
			Init,
			Sync,
			Gen,
			NewCmd,
		},
	}

	return app.Run(os.Args)
}

func flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "tmpl",
			Aliases:     []string{"t"},
			Value:       "temp",
			Usage:       "template root dir",
			Destination: &opts.TmplDir,
		},
		&cli.StringFlag{
			Name:        "conf",
			Aliases:     []string{"f"},
			Value:       "",
			Usage:       "config file path",
			Destination: &opts.Config,
		},
		&cli.StringFlag{
			Name:        "db-driver",
			Aliases:     []string{"dd"},
			Value:       "postgres",
			Usage:       "database driver name: mysql|postgres",
			Destination: &opts.DBDriver,
		},
		&cli.StringFlag{
			Name:        "db-name",
			Aliases:     []string{"dn"},
			Value:       "Ardan",
			Usage:       "database name",
			Destination: &opts.DBName,
		},
		&cli.StringFlag{
			Name:        "db-conn",
			Aliases:     []string{"dc"},
			Usage:       "database connection string",
			Destination: &opts.DBConnStr,
		},
		&cli.StringFlag{
			Name:        "mod",
			Aliases:     []string{"m"},
			Usage:       "go module name",
			Required:    false,
			Destination: &opts.GoModName,
		},
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Value:       "./output",
			Usage:       "output directory",
			Destination: &opts.OutputDir,
		},
		&cli.BoolFlag{
			Name:        "sample",
			Aliases:     []string{"s"},
			Value:       false,
			Usage:       "initial sample code",
			Destination: &opts.Sample,
		},
		// project layout
		&cli.StringFlag{
			Name:        "cmd",
			Value:       "cmd",
			Usage:       "[Layout] executed command root directory",
			Destination: &opts.CmdDir,
		},
		&cli.StringFlag{
			Name:        "doc",
			Value:       "doc",
			Usage:       "[Layout] documents root directory",
			Destination: &opts.DocDir,
		},
		&cli.StringFlag{
			Name:        "app",
			Value:       "app",
			Usage:       "[Layout] application layer root directory",
			Destination: &opts.AppDir,
		},
		&cli.StringFlag{
			Name:        "model",
			Value:       "model",
			Usage:       "[Layout] application layer domain directory",
			Destination: &opts.ModelDir,
		},
		&cli.StringFlag{
			Name:        "service",
			Value:       "service",
			Usage:       "[Layout] application layer service directory",
			Destination: &opts.ServiceDir,
		},
		&cli.StringFlag{
			Name:        "repository",
			Value:       "repository",
			Usage:       "[Layout] application layer repository directory",
			Destination: &opts.RepositoryDir,
		},
		&cli.StringFlag{
			Name:        "server",
			Value:       "server",
			Usage:       "[Layout] server layer root directory",
			Destination: &opts.ServerDir,
		},
		&cli.StringFlag{
			Name:        "module",
			Value:       "module",
			Usage:       "[Layout] server layer module directory",
			Destination: &opts.ServerModuleDir,
		},
		&cli.StringFlag{
			Name:        "global",
			Value:       "global",
			Usage:       "[Layout] server layer global directory",
			Destination: &opts.ServerGlobalDir,
		},
		&cli.StringFlag{
			Name:        "controller",
			Value:       "controller",
			Usage:       "[Layout] server layer controller directory",
			Destination: &opts.ControllerDir,
		},
		&cli.StringFlag{
			Name:        "middleware",
			Value:       "middleware",
			Usage:       "[Layout] server layer middleware directory",
			Destination: &opts.MiddlewareDir,
		},
		&cli.StringFlag{
			Name:        "handler",
			Value:       "handler",
			Usage:       "[Layout] server layer handler directory",
			Destination: &opts.HandlerDir,
		},
	}
}

func bashComplete(c *cli.Context) {
	// fmt.Fprintf(c.App.Writer, "bashComplete\n")
}

func before(c *cli.Context) error {
	// fmt.Fprintf(c.App.Writer, "before\n")
	var err error
	err = Setup(c)
	return err
}

func after(c *cli.Context) error {
	return nil
}
