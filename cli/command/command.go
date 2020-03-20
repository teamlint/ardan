package command

import (
	"fmt"
	"sync"

	"github.com/teamlint/ardan/cli/setting"
	"github.com/urfave/cli/v2"
)

var (
	once    sync.Once
	Setting *setting.Setting
)

func Setup(c *cli.Context) error {
	// check go mod
	if GoModName == "" {
		fmt.Fprintf(c.App.Writer, "please use `go mod init` or use `--mod`\n")
	}
	cli.Exit(":(", -1)

	// init setting
	once.Do(func() {
		Setting = setting.New(setting.Options{
			TmplDir:   TemplateDir,
			OutputDir: OutputDir,
			Sample:    Sample,
			DBDriver:  DBDriver,
			DBName:    DBName,
			DBConnStr: DBConnStr,
			GoModName: GoModName,
			// layout
			CmdDir:          CmdDir,
			DocDir:          DocDir,
			AppDir:          AppDir,
			ModelDir:        ModelDir,
			ServiceDir:      ServiceDir,
			RepositoryDir:   RepositoryDir,
			ServerDir:       ServerDir,
			ServerModuleDir: ServerModuleDir,
			ServerGlobalDir: ServerGlobalDir,
			ControllerDir:   ControllerDir,
			HandlerDir:      HandlerDir,
			MiddlewareDir:   MiddlewareDir,
		})
	})
	info()
	return nil
}
func info() {
	fmt.Printf("templates=%v\n", TemplateDir)
	fmt.Printf("config=%v\n", ConfigFile)
	fmt.Printf("db-driver=%v\n", DBDriver)
	fmt.Printf("db-name=%v\n", DBName)
	fmt.Printf("db-conn=%v\n", DBConnStr)
	fmt.Printf("goPkgName=%v\n", GoPkgName)
	fmt.Printf("goModName=%v\n", GoModName)
	fmt.Printf("output=%v\n", OutputDir)
	fmt.Printf("sample=%v\n", Sample)
}
