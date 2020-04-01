package command

import (
	"fmt"
	"path/filepath"

	"github.com/teamlint/ardan/cli/setting"
	"github.com/teamlint/ardan/config"
	"github.com/teamlint/ardan/pkg"
	"github.com/urfave/cli/v2"
)

// Sync sync database tabels struct
var Sync = &cli.Command{
	Name:    "sync",
	Aliases: []string{"s"},
	Usage:   "sync to database",
	Action: func(c *cli.Context) error {
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
	if Setting.DBConnStr == "" {
		// read config.yml
		configFile := filepath.Join(Setting.Output, Setting.Cmd, Setting.Server, setting.ConfigFile)
		if pkg.Exists(configFile) {
			if err := config.LoadFile(configFile); err != nil {
				return fmt.Errorf("read config file err=%v\n", err)
			}
			connStr := config.Get("Databases", Setting.DBName, "ConnString").String("")
			if connStr == "" {
				return ErrDBConnStrNone
			}
			Setting.DBConnStr = connStr
			info(c, "found `db-conn` in %v [%v] section\n", configFile, fmt.Sprintf("Databases.%s.ConnString", Setting.DBName))
		} else {
			return ErrDBConnStrNone
		}
	}
	info(c, "db-conn = %v\n", Setting.DBConnStr)
	// gen main.go
	if err := genMainfile(c); err != nil {
		return err
	}
	// run main
	info(c, ">> %v\n", "model sync done")

	return nil
}

func genMainfile(c *cli.Context) error {
	beans, err := ParseModelFiles(c, setting.DirectiveSync)
	if err != nil {
		return err
	}
	name := Setting.TmplFile(Setting.Cmd, Setting.Sync, setting.GoMainFileTmpl)
	target := Setting.TargetFile(name)
	err = Render(target, name, map[string]interface{}{"Models": beans})
	if err != nil {
		return fmt.Errorf("sync.Render err=%v\n", err)
	}
	// go run
	output, err := GoRun(filepath.Dir(target), setting.GoMainFile)
	if err != nil {
		return fmt.Errorf("sync.GoRun err=%v\n", err)
	}
	info(c, "--------------------- [sync] start ----------------------------\n"+output)
	pkg.DeleteFile(target, true)
	info(c, "---------------------- [sync] end -----------------------------\n")

	return nil
}
