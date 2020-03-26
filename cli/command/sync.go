package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/teamlint/go-astra"
	"github.com/teamlint/go-astra/types"
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
	// path := "./output/app/model/demo.go"
	// file, err := astra.ParseFile(path)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// t, err := json.Marshal(file)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(t))

	// gen main.go
	if err := genMainfile(c); err != nil {
		return err
	}
	// run main

	info(c, ">> %v\n", "model sync done")

	return nil
}

func parseFiles(c *cli.Context) ([]string, error) {
	beans := make([]string, 0)
	info(c, "sync.parseFiles starting...\n")
	// for _, name := range Setting.Codes {
	root := filepath.Join(Setting.Output, Setting.App, Setting.Model)
	err := filepath.Walk(root, func(path string, fi os.FileInfo, e1 error) error {
		if fi.IsDir() {
			return nil
		}
		info(c, "sync.parseFiles path=%v\n", path)
		// target := Setting.TargetFile(name)
		// info(c, "sync.parseFiles found model file=%v\n", target)
		tf, err := astra.ParseFile(path)
		if err != nil {
			return fmt.Errorf("sync.parseFiles err=%v\n", err)
		}
		// t, err := json.Marshal(tf)
		// if err != nil {
		// 	log.Println(err)
		// }
		// log.Println(string(t))
		beans = append(beans, parseSyncModel(c, tf)...)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("sync.walk err=%v\n", err)
	}
	info(c, "sync.parseFiles completed.\n")
	return beans, nil
}

func parseSyncModel(c *cli.Context, tf *types.File) []string {
	beans := make([]string, 0)
	// for _, tf := range tfs {
	for _, m := range tf.Structures {
		info(c, "sync.parseSyncModel model.Docs=%v\n", m.Docs)
		for _, doc := range m.Docs {
			if strings.HasPrefix(doc, "//ardan:model sync") {
				info(c, "sync.parseSyncModel found model=%v\n", m.Name)
				beans = append(beans, m.Name)
			}
		}
	}
	// }
	return beans
}

func genMainfile(c *cli.Context) error {
	// func GenMainfile(binaryName, path string, tfs []*types.File) error {
	// log.Println("Creating mainfile at", path)

	// f, err := pkg.NewFile(path)
	// if err != nil {
	// 	return fmt.Errorf("error creating generated mainfile: %v", err)
	// }
	// defer f.Close()

	// if err := mainfileTemplate.Execute(f, data); err != nil {
	// 	return fmt.Errorf("can't execute mainfile template: %v", err)
	// }
	// if err := f.Close(); err != nil {
	// 	return fmt.Errorf("error closing generated mainfile: %v", err)
	// }
	// // we set an old modtime on the generated mainfile so that the go tool
	// // won't think it has changed more recently than the compiled binary.
	// longAgo := time.Now().Add(-time.Hour * 24 * 365 * 10)
	// if err := os.Chtimes(path, longAgo, longAgo); err != nil {
	// 	return fmt.Errorf("error setting old modtime on generated mainfile: %v", err)
	// }
	// return nil

	beans, err := parseFiles(c)
	if err != nil {
		return err
	}
	info(c, "beans=%v\n", beans)
	name := filepath.Join("/", Setting.Cmd, "sync", "main.go.tmpl")
	target := Setting.TargetFile(name)
	err = Render(target, name, map[string]interface{}{"Beans": beans})
	if err != nil {
		return fmt.Errorf("sync.Render err=%v\n", err)
	}

	return nil
}
