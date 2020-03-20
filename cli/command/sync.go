package command

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/teamlint/go-astra"
	"github.com/urfave/cli/v2"
)

// Sync sync database tabels struct
var Sync = &cli.Command{
	Name:    "sync",
	Aliases: []string{"s"},
	Usage:   "sync to database",
	// Flags: []cli.Flag{
	// 	&cli.StringFlag{
	// 		Name:        "db",
	// 		Aliases:     []string{"d"},
	// 		Value:       "postgres",
	// 		Usage:       "database name: mysql|postgres",
	// 		Destination: &DB,
	// 	},
	// 	&cli.StringFlag{
	// 		Name:        "conn",
	// 		Aliases:     []string{"c"},
	// 		Value:       "ardan",
	// 		Usage:       "database connection string",
	// 		Destination: &ConnStr,
	// 	},
	// },
	Action: func(c *cli.Context) error {
		fmt.Println("sync root command")
		var err error
		err = syncToDB(c)
		if err != nil {
			return err
		}
		return nil
	},
	Subcommands: []*cli.Command{
		InitApp,
	},
}

// SyncToDB
var SyncToDB = &cli.Command{
	Name:   "todb",
	Usage:  "sync domain model to database",
	Action: syncToDB,
}

func syncToDB(c *cli.Context) error {
	// log.Printf("template engine=%+v appDir=%v\n", *set.Template, set.AppDir)
	path := filepath.Join(TemplateDir, "/app/model/user.go")
	file, err := astra.ParseFile(path)
	if err != nil {
		fmt.Println(err)
	}
	t, err := json.Marshal(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(t))

	return nil
}
