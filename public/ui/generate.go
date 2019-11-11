// +build ignore

package main

import (
	"log"

	"github.com/sapk/rtme-browser/public/ui"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(ui.WebApp, vfsgen.Options{
		PackageName:  "ui",
		VariableName: "WebApp",
		BuildTags:    "!dev",
		//Filename:     "webapp.go",
	})
	if err != nil {
		log.Fatal("%v", err)
	}
}
