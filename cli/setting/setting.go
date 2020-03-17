package setting

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/teamlint/ardan/cli/lib"
)

var (
	once     sync.Once
	instance *Setting
)

const (
	TmplExt = ".tmpl" // template file extention
)

type Setting struct {
	Template  *template.Template
	AppDir    string // application dir
	CmdDir    string // cmd dir
	CtrlDir   string // controller dir
	MdwDir    string // middleware dir
	ServerDir string // server dir
	ModuleDir string // server dir
	// ----
	templateDir string
	outputDir   string
}

// Init init settings
func Init(tmplDir string, outputDir string) {
	if !lib.Exists(tmplDir) {
		msg := "template dir is not exists"
		log.Fatal(msg)
		panic(msg)
	}
	tmpl := template.New("ardan")
	// _, err := tmpl.ParseGlob(filepath.Join(tmplDir, "*/*"+TmplExt))
	_, err := findAndParseTemplates(tmplDir, template.FuncMap{})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	appDir := filepath.Join(outputDir, "app")
	instance = &Setting{
		Template: tmpl,
		AppDir:   appDir,
	}

}

// Instance
func Instance() *Setting {
	return instance
}
func findAndParseTemplates(rootDir string, funcMap template.FuncMap) (*template.Template, error) {
	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, TmplExt) {
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]
			t := root.New(name).Funcs(funcMap)
			t, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}

		return nil
	})

	return root, err
}
