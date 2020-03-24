package command

import (
	"encoding/json"
	"fmt"

	"github.com/teamlint/go-astra"
	"github.com/urfave/cli/v2"
)

// Sync sync database tabels struct
var Sync = &cli.Command{
	Name:    "sync",
	Aliases: []string{"s"},
	Usage:   "sync to database",
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
		SyncToDB,
	},
}

// SyncToDB
var SyncToDB = &cli.Command{
	Name:   "todb",
	Usage:  "sync domain model to database",
	Action: syncToDB,
}

func syncToDB(c *cli.Context) error {
	// path := filepath.Join(Setting.Output, Setting.App, Setting.Model, "demo_user.go")
	path := "./go.mod"
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
