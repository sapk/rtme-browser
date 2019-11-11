// +build ignore

package ui

import (
	"log"

	"gitea.com/sapk/explore/public/webapp"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(webapp.WebApp, vfsgen.Options{
		PackageName:  "ui",
		VariableName: "WebApp",
		BuildTags:    "!dev",
		//Filename:     "webapp.go",
	})
	if err != nil {
		log.Fatal("%v", err)
	}
}
