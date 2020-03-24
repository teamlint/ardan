package setting

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/teamlint/ardan/pkg"
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
	GoMod     string
	// layout
	Output       string // output root dir
	Cmd          string // cmd dir
	Doc          string // documents root directory
	App          string // application dir
	Model        string // domain layer directory
	Service      string // service layer directory
	Repository   string // repository layer directory
	Server       string // server layer directory
	ServerModule string // server module directory
	ServerGlobal string // server global directory
	Controller   string // controller directory
	Handler      string // handler directory
	Middleware   string // middleware directory
	Sample       bool
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
	Config    string // config file
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
	if !pkg.Exists(opt.TmplDir) {
		msg := "template dir is not exists"
		log.Fatal(msg)
		panic(msg)
	}

	outputDir := clean(opt.OutputDir)
	cmdDir := clean(opt.CmdDir)
	docDir := clean(opt.DocDir)
	appDir := clean(opt.AppDir)
	modelDir := clean(opt.ModelDir)
	serviceDir := clean(opt.ServiceDir)
	repositoryDir := clean(opt.RepositoryDir)
	serverDir := clean(opt.ServerDir)
	serverModuleDir := clean(opt.ServerModuleDir)
	serverGlobalDir := clean(opt.ServerGlobalDir)
	controllerDir := clean(opt.ControllerDir)
	handlerDir := clean(opt.HandlerDir)
	middlewareDir := clean(opt.MiddlewareDir)

	instance := &Setting{
		Layouts:   make([]string, 0),
		Codes:     make([]string, 0),
		Samples:   make([]string, 0),
		DBDriver:  opt.DBDriver,
		DBName:    opt.DBName,
		DBConnStr: opt.DBConnStr,
		GoMod:     opt.GoModName,
		// layout
		Output:       outputDir,
		Cmd:          cmdDir,
		Doc:          docDir,
		App:          appDir,
		Model:        modelDir,
		Service:      serviceDir,
		Repository:   repositoryDir,
		Server:       serverDir,
		ServerModule: serverModuleDir,
		ServerGlobal: serverGlobalDir,
		Controller:   controllerDir,
		Handler:      handlerDir,
		Middleware:   middlewareDir,
		Sample:       opt.Sample,
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
	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if len(path) < pfx {
			return nil
		}
		name := path[pfx:]
		// log.Printf("path=%v\n", path)
		// log.Printf("path_name=%v\n", name)
		// is dir, make it
		if info.IsDir() {
			// if err := pkg.Mkdir(filepath.Join(opt.OutputDir, name)); err != nil {
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

			b, e2 := pkg.GetFileContent(path)
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
			b, e2 := pkg.GetFileContent(path)
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
	fm["clean"] = clean
	fm["randomString"] = pkg.RandomString
	return fm
}

func clean(path string) string {
	path = strings.TrimPrefix(path, ".")
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")
	return path
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
		return filepath.Join(s.Output, dst)
		// default:
		// 	log.Printf("[TargetFile].default ext=%v\n", ext)
	}
	return filepath.Join(s.Output, srcname)
}
