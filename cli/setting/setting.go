package setting

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/teamlint/ardan/cli/lib"
)

type TmplType = string

const (
	TmplTypeCode   TmplType = ".tmpl"
	TmplTypeView   TmplType = ".tmpl"
	TmplTypeSample TmplType = ".samp"
	TmplTypeBuild  TmplType = ".ardan"
)

type Directive = string

const (
	DirectiveSync Directive = "ardan:sync"
)

type Setting struct {
	Template  *template.Template
	Layouts   []string // project layout dir names
	Codes     []string // source code names
	Samples   []string // sample names
	DBDriver  string   // database driver name
	DBName    string   // database name
	DBConnStr string   // database connection string
	GoModName string
	// layout
	OutputDir       string // output root dir
	CmdDir          string // cmd dir
	DocDir          string // documents root directory
	AppDir          string // application dir
	ModelDir        string // domain layer directory
	ServiceDir      string // service layer directory
	RepositoryDir   string // repository layer directory
	ServerDir       string // server layer directory
	ServerModuleDir string // server module directory
	ServerGlobalDir string // server global directory
	ControllerDir   string // controller directory
	HandlerDir      string // handler directory
	MiddlewareDir   string // middleware directory
	Sample          bool
}

// Options setting options
type Options struct {
	TmplDir   string
	OutputDir string
	FuncMap   template.FuncMap
	DBDriver  string // database driver name
	DBName    string // database name
	DBConnStr string // database connection string
	GoModName string
	// project layout
	CmdDir          string // command root directory
	DocDir          string // documents root directory
	AppDir          string // application layer directory
	ModelDir        string // domain layer directory
	ServiceDir      string // service layer directory
	RepositoryDir   string // repository layer directory
	ServerDir       string // server layer directory
	ServerModuleDir string // server module directory
	ServerGlobalDir string // server global directory
	ControllerDir   string // controller directory
	HandlerDir      string // handler directory
	MiddlewareDir   string // middleware directory
	Sample          bool
}

// New init settings
func New(opt Options) *Setting {
	if !lib.Exists(opt.TmplDir) {
		msg := "template dir is not exists"
		log.Fatal(msg)
		panic(msg)
	}

	// cmdDir := filepath.Join(opt.OutputDir, opt.CmdDir)
	// docDir := filepath.Join(opt.OutputDir, opt.DocDir)
	// appDir := filepath.Join(opt.OutputDir, opt.AppDir)
	// modelDir := filepath.Join(opt.OutputDir, opt.AppDir, opt.ModelDir)
	// serviceDir := filepath.Join(opt.OutputDir, opt.AppDir, opt.ServiceDir)
	// repositoryDir := filepath.Join(opt.OutputDir, opt.AppDir, opt.RepositoryDir)
	// serverDir := filepath.Join(opt.OutputDir, opt.ServerDir)
	// serverModuleDir := filepath.Join(opt.OutputDir, opt.ServerDir, opt.ServerModuleDir)
	// serverGlobalDir := filepath.Join(opt.OutputDir, opt.ServerDir, opt.ServerGlobalDir)
	// controllerDir := filepath.Join(opt.OutputDir, opt.ServerDir, opt.ControllerDir)
	// handlerDir := filepath.Join(opt.OutputDir, opt.ServerDir, opt.HandlerDir)
	// middlewareDir := filepath.Join(opt.OutputDir, opt.ServerDir, opt.MiddlewareDir)
	cmdDir := opt.CmdDir
	docDir := opt.DocDir
	appDir := opt.AppDir
	modelDir := opt.ModelDir
	serviceDir := opt.ServiceDir
	repositoryDir := opt.RepositoryDir
	serverDir := opt.ServerDir
	serverModuleDir := opt.ServerModuleDir
	serverGlobalDir := opt.ServerGlobalDir
	controllerDir := opt.ControllerDir
	handlerDir := opt.HandlerDir
	middlewareDir := opt.MiddlewareDir

	instance := &Setting{
		Layouts:   make([]string, 0),
		Codes:     make([]string, 0),
		Samples:   make([]string, 0),
		DBDriver:  opt.DBDriver,
		DBName:    opt.DBName,
		DBConnStr: opt.DBConnStr,
		GoModName: opt.GoModName,
		// layout
		OutputDir:       opt.OutputDir,
		CmdDir:          cmdDir,
		DocDir:          docDir,
		AppDir:          appDir,
		ModelDir:        modelDir,
		ServiceDir:      serviceDir,
		RepositoryDir:   repositoryDir,
		ServerDir:       serverDir,
		ServerModuleDir: serverModuleDir,
		ServerGlobalDir: serverGlobalDir,
		ControllerDir:   controllerDir,
		HandlerDir:      handlerDir,
		MiddlewareDir:   middlewareDir,
		Sample:          opt.Sample,
	}

	err := instance.walkTemplates(opt)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return instance

}

func (s *Setting) walkTemplates(opt Options) error {
	cleanRoot := filepath.Clean(opt.TmplDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	// log.Printf("root=%v\n", cleanRoot)
	log.Printf("pfx=%v\n", pfx)
	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if len(path) < pfx {
			return nil
		}
		name := path[pfx:]
		// log.Printf("path=%v\n", path)
		// log.Printf("path_name=%v\n", name)
		// is dir, make it
		if info.IsDir() {
			// if err := lib.Mkdir(filepath.Join(opt.OutputDir, name)); err != nil {
			// 	return err
			// }
			s.Layouts = append(s.Layouts, filepath.Join(opt.OutputDir, name))
			return nil
		}

		// // is template
		if strings.HasSuffix(path, TmplTypeCode) {
			if e1 != nil {
				return e1
			}

			b, e2 := lib.GetFileContent(path)
			if e2 != nil {
				return e2
			}
			// log.Printf("temp_name=%v\n", name)
			t := root.New(name).Funcs(defaultFuncMap()).Funcs(opt.FuncMap)
			t, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
			s.Codes = append(s.Codes, name)
			return nil
		}
		// is sample
		if opt.Sample && strings.HasSuffix(path, TmplTypeSample) {
			if e1 != nil {
				return e1
			}
			b, e2 := lib.GetFileContent(path)
			if e2 != nil {
				return e2
			}
			t := root.New(name).Funcs(defaultFuncMap()).Funcs(opt.FuncMap)
			t, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
			s.Samples = append(s.Samples, name)
			return nil
		}

		return nil
	})

	s.Template = root
	return err
}

func defaultFuncMap() template.FuncMap {
	fm := template.FuncMap{}
	fm["clean"] = func(path string) string {
		path = strings.TrimPrefix(path, ".")
		path = strings.TrimPrefix(path, "/")
		path = strings.TrimSuffix(path, "/")
		return path
	}
	return fm
}

// findDirective return the first line of a doc which contains a directive
// the directive and '//' are removed
func findDirective(doc []string, directive string) (string, bool) {
	if len(doc) < 1 {
		return "", false
	}

	// check lines of doc for directive
	for _, c := range doc {
		t := strings.TrimLeft(c, "/")
		if strings.HasPrefix(t, directive) {
			return c, true
		}
	}

	return "", false
}
func (s *Setting) TargetFile(srcname string) string {
	ext := filepath.Ext(srcname)

	switch ext {
	case TmplTypeCode, TmplTypeSample, TmplTypeBuild:
		dst := strings.TrimSuffix(srcname, ext)
		// log.Printf("[TargetFile] dst=%v,ext=%v\n", dst, filepath.Ext(dst))
		return filepath.Join(s.OutputDir, dst)
		// default:
		// 	log.Printf("[TargetFile].default ext=%v\n", ext)
	}
	return filepath.Join(s.OutputDir, srcname)
}
