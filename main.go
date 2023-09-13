package main

import (
	"fmt"
	"github.com/lprashanthgowda/kube-crd2openapi/pkg/cmd"
	"os"
)

func main() {
	app := cmd.NewRootCommand()
	err := app.Execute()
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
	}
}
