package command

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	execute "github.com/alexellis/go-execute/pkg/v1"
	"github.com/teamlint/ardan/cli/setting"
	"github.com/teamlint/ardan/pkg"
	"github.com/urfave/cli/v2"
	"github.com/vetcher/go-astra"
	"github.com/vetcher/go-astra/types"
)

const (
	LogPrefix      = "[ardan] "
	RepositoryName = "Repository"
	ServiceName    = "Service"
	ControllerName = "Controller"
)

type Model struct {
	Name      string
	Directive string
	Gen       setting.GenSet
	Struct    types.Struct
}

var (
	once             sync.Once
	Setting          *setting.Setting
	ErrGoModNone     = errors.New("please use `go mod init` or use `--mod` global options\n")
	ErrDBConnStrNone = errors.New("please use `--db-conn` global options setting database connection string\n")
)

func Setup(c *cli.Context) error {
	// check go mod
	if opts.GoModName == "" {
		gomod := readGoMod()
		if gomod == "" {
			return ErrGoModNone
		}
		// info(c, "go.mod name = %v\n", gomod)
		opts.GoModName = gomod
	}

	// init setting
	once.Do(func() {
		Setting = setting.New(opts)
	})

	// info(c, "setting init done.\n")
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

func GoBuild(dir string, source string, name string) error {
	return nil
}

func GoRun(dir string, main string) (string, error) {
	cmd := execute.ExecTask{
		Command: "go",
		// Args:         []string{"run", "-ldflags", `"-w -s"`, "-o", name, source},
		Args:         []string{"run", main},
		PrintCommand: false, // print command
		Cwd:          dir,
		// - GOOS=linux go build -ldflags '-w -s' -o ./release/{{.Product}} ./cmd/server/main.go
	}
	result, err := cmd.Execute()
	if err != nil {
		return "", fmt.Errorf("GoRun err=%v\n", err)
	}
	if result.ExitCode != 0 {
		return "", fmt.Errorf("GoRun err=%v, ExitCode=%v\n", err, result.ExitCode)
	}
	return result.Stdout, nil

}

func info(c *cli.Context, msg string, args ...interface{}) {
	fmt.Fprintf(c.App.Writer, LogPrefix+msg, args...)
}

func Render(fname string, tmplName string, data map[string]interface{}) error {
	var buf bytes.Buffer
	err := Setting.Template.ExecuteTemplate(&buf, tmplName, tmplData(data))
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

func tmplData(data map[string]interface{}) map[string]interface{} {
	main := make(map[string]interface{})
	main["Setting"] = Setting
	if data != nil && len(data) > 0 {
		for k, v := range data {
			main[k] = v
		}
	}
	return main
}

func ParseModelFiles(c *cli.Context, directive setting.Directive) ([]*Model, error) {
	beans := make([]*Model, 0)
	info(c, "ParseModelFiles starting...\n")
	root := filepath.Join(Setting.Output, Setting.App, Setting.Model)
	err := filepath.Walk(root, func(path string, fi os.FileInfo, e1 error) error {
		if fi.IsDir() {
			return nil
		}
		// info(c, "\tParseModelFiles path=%v\n", path)
		tf, err := astra.ParseFile(path)
		if err != nil {
			return fmt.Errorf("ParseModelFiles err=%v\n", err)
		}
		beans = append(beans, parseModelDiretive(c, tf, directive)...)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("walk model err=%v\n", err)
	}
	info(c, "ParseModelFiles completed.\n")
	return beans, nil
}

func parseModelDiretive(c *cli.Context, tf *types.File, directive setting.Directive) []*Model {
	models := make([]*Model, 0)
	for _, m := range tf.Structures {
		// info(c, "\tparseSyncModel model.Docs=%v, model=%v\n", m.Docs, m)
		for _, doc := range m.Docs {
			if dire, ok := Setting.HasDirective(doc, directive); ok {
				model := Model{Name: m.Name, Directive: dire, Struct: m}
				switch dire {
				case setting.DirectiveGen:
					// gen direction
					gs, err := Setting.ParseDirectiveGen(doc)
					if err != nil {
						log.Fatal(err)
					}
					// default names
					if gs.Repository == "" {
						gs.Repository = m.Name + RepositoryName
					}
					if gs.Service == "" {
						gs.ServiceInterface = m.Name + ServiceName
						gs.Service = pkg.LowerFirst(m.Name) + ServiceName
					}
					if gs.Controller == "" {
						gs.Controller = m.Name + ControllerName
					}
					model.Gen = *gs
				case setting.DirectiveSync:
					// other direction
				}
				models = append(models, &model)
				info(c, "found gen.model=%v, directive=%v, repository=%v, service=%v:%v, controller=%v\n", model.Name, model.Directive, model.Gen.Repository, model.Gen.Service, model.Gen.ServiceInterface, model.Gen.Controller)
			}
		}
	}
	return models
}
