package setting

import (
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
	// template file extention name
	TmplExt = ".tmpl" // template file extention
	DemoExt = ".demo" // demo file extention
)

type Setting struct {
	Template *template.Template
	// layout
	CmdDir          string // cmd dir
	DocDir          string // documents root directory
	AppDir          string // application dir
	DomainDir       string // domain layer directory
	ServiceDir      string // service layer directory
	RepositoryDir   string // repository layer directory
	ServerDir       string // server layer directory
	ServerModuleDir string // server module directory
	ServerGlobalDir string // server global directory
	ControllerDir   string // controller directory
	HandlerDir      string // handler directory
	MiddlewareDir   string // middleware directory
	Demo            bool
}

// Options setting options
type Options struct {
	TmplDir   string
	OutputDir string
	FuncMap   template.FuncMap
	// project layout
	CmdDir          string // command root directory
	DocDir          string // documents root directory
	AppDir          string // application layer directory
	DomainDir       string // domain layer directory
	ServiceDir      string // service layer directory
	RepositoryDir   string // repository layer directory
	ServerDir       string // server layer directory
	ServerModuleDir string // server module directory
	ServerGlobalDir string // server global directory
	ControllerDir   string // controller directory
	HandlerDir      string // handler directory
	MiddlewareDir   string // middleware directory
	Demo            bool
}

// Init init settings
func Init(opt Options) {
	if !lib.Exists(opt.TmplDir) {
		msg := "template dir is not exists"
		log.Fatal(msg)
		panic(msg)
	}
	// tmpl := template.New("ardan")
	// _, err := tmpl.ParseGlob(filepath.Join(tmplDir, "*/*"+TmplExt))
	tmpl, err := walkTemplates(opt)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	cmdDir := filepath.Join(opt.OutputDir, opt.CmdDir)
	docDir := filepath.Join(opt.OutputDir, opt.DocDir)
	appDir := filepath.Join(opt.OutputDir, opt.AppDir)
	domainDir := filepath.Join(opt.OutputDir, opt.AppDir, opt.DomainDir)
	serviceDir := filepath.Join(opt.OutputDir, opt.AppDir, opt.ServiceDir)
	repositoryDir := filepath.Join(opt.OutputDir, opt.AppDir, opt.RepositoryDir)
	serverDir := filepath.Join(opt.OutputDir, opt.ServerDir)
	serverModuleDir := filepath.Join(opt.OutputDir, opt.ServerDir, opt.ServerModuleDir)
	serverGlobalDir := filepath.Join(opt.OutputDir, opt.ServerDir, opt.ServerGlobalDir)
	controllerDir := filepath.Join(opt.OutputDir, opt.ServerDir, opt.ControllerDir)
	handlerDir := filepath.Join(opt.OutputDir, opt.ServerDir, opt.HandlerDir)
	middlewareDir := filepath.Join(opt.OutputDir, opt.ServerDir, opt.MiddlewareDir)

	instance = &Setting{
		Template: tmpl,
		// layout
		CmdDir:          cmdDir,
		DocDir:          docDir,
		AppDir:          appDir,
		DomainDir:       domainDir,
		ServiceDir:      serviceDir,
		RepositoryDir:   repositoryDir,
		ServerDir:       serverDir,
		ServerModuleDir: serverModuleDir,
		ServerGlobalDir: serverGlobalDir,
		ControllerDir:   controllerDir,
		HandlerDir:      handlerDir,
		MiddlewareDir:   middlewareDir,
		Demo:            opt.Demo,
	}

}

// Instance
func Instance() *Setting {
	return instance
}
func walkTemplates(opt Options) (*template.Template, error) {
	cleanRoot := filepath.Clean(opt.TmplDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	// log.Printf("root=%v\n", cleanRoot)
	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if len(path) < pfx {
			return nil
		}
		name := path[pfx:]
		// log.Printf("path=%v\n", path)
		// log.Printf("path_name=%v\n", name)
		// is dir, make it
		if info.IsDir() {
			if err := lib.Mkdir(filepath.Join(opt.OutputDir, name)); err != nil {
				return err
			}
			return nil
		}
		// is template
		if strings.HasSuffix(path, TmplExt) {
			if e1 != nil {
				return e1
			}

			b, e2 := lib.GetFileContent(path)
			if e2 != nil {
				return e2
			}

			name := path
			// log.Printf("temp_name=%v\n", name)
			t := root.New(name).Funcs(defaultFuncMap()).Funcs(opt.FuncMap)
			t, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}
		// is demo
		if opt.Demo && strings.HasSuffix(path, DemoExt) {
			if e1 != nil {
				return e1
			}
			dest := filepath.Join(opt.OutputDir, strings.TrimSuffix(name, DemoExt))
			log.Printf("demo.src=%v, demo.dest=%v\n", path, dest)
			e2 := lib.Copy(path, dest)
			if e2 != nil {
				return e2
			}
		}

		return nil
	})

	return root, err
}

func defaultFuncMap() template.FuncMap {
	return template.FuncMap{}
}
