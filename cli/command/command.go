package command

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"log"
	"os/exec"
	"strings"
	"sync"

	execute "github.com/alexellis/go-execute/pkg/v1"
	"github.com/teamlint/ardan/cli/setting"
	"github.com/teamlint/ardan/pkg"
	"github.com/urfave/cli/v2"
)

const (
	LogPrefix = "[ardan] "
)

var (
	once         sync.Once
	Setting      *setting.Setting
	ErrGoModNone = errors.New("please use `go mod init` or use `--mod` global options\n")
)

func Setup(c *cli.Context) error {
	// check go mod
	if opts.GoModName == "" {
		gomod := readGoMod()
		if gomod == "" {
			return ErrGoModNone
		}
		info(c, "go.mod name = %v\n", gomod)
		opts.GoModName = gomod
	}
	// cli.Exit(":(", -1)

	// init setting
	once.Do(func() {
		Setting = setting.New(opts)
	})

	info(c, "setting init done.\n")
	return nil
}

func readGoMod() string {
	cmd := execute.ExecTask{
		Command: "go",
		Args:    []string{"list", "-m"},
	}
	result, err := cmd.Execute()
	if err != nil {
		return ""
	}
	if result.ExitCode != 0 {
		return ""
	}
	res := strings.Trim(result.Stdout, "\n")

	return res

}

func info(c *cli.Context, msg string, args ...interface{}) {
	fmt.Fprintf(c.App.Writer, LogPrefix+msg, args...)
}

func Render(fname string, tmplName string) error {
	var buf bytes.Buffer
	err := Setting.Template.ExecuteTemplate(&buf, tmplName, tmplData())
	if err != nil {
		log.Printf("template[%v] render error=%v\n", tmplName, err)
		return err
	}
	// format
	if strings.HasSuffix(fname, ".go") {
		data, err := format.Source(buf.Bytes())
		if err != nil {
			log.Printf("template[%v] format source error=%v\n", tmplName, err)
			return err
		}
		// log.Printf("format.source=%v\n", string(data))
		err = pkg.WriteFile(fname, data)
		if err != nil {
			log.Printf("template[%v] write file[%v] error=%v\n", tmplName, fname, err)
			return err
		}
		defer func() {
			err := exec.Command("goimports", "-w", fname).Run()
			if err != nil {
				log.Printf("`goimports -w %v` error=%v\n", fname, err)
			}
			exec.Command("gofmt", "-w", "-s", fname).Run()
			if err != nil {
				log.Printf("`gofmt -w -s %v` error=%v\n", fname, err)
			}
		}()
	} else {
		err = pkg.WriteFile(fname, buf.Bytes())
		if err != nil {
			log.Printf("template[%v] write file[%v] error=%v\n", tmplName, fname, err)
			return err
		}
	}
	return nil
}

func tmplData() map[string]interface{} {
	data := make(map[string]interface{})
	data["Setting"] = Setting
	return data
}
