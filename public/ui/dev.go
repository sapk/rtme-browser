// +build dev

package ui

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	_, d := filepath.Split(pwd)
	if d == "ui" { //For go generate command
		WebApp = http.Dir("../../assets/ui/dist")
	}
}

// WebApp contains project WebApp assets.
var WebApp http.FileSystem = http.Dir("assets/ui/dist")
