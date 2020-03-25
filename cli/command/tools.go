// +build tools

package command

import (
	_ "github.com/go-task/task/cmd/task"
	_ "golang.org/x/tools/cmd/goimports"
	_ "golang.org/x/tools/cmd/stringer"
)
