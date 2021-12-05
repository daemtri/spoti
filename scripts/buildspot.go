//go:build spot

package main

import (
	"fmt"
	"github.com/bitfield/script"
	"log"
	"runtime"
)

func main() {
	appName := "spotexample-calc"
	version := must(script.Exec(`git describe --abbrev=0 --tags`).String())
	commit := must(script.Exec(`git log --pretty=format:"%h" -1`).String())
	branch := must(script.Exec(`git rev-parse --abbrev-ref HEAD`).String())
	distPath := fmt.Sprintf("dist/%s", appName)
	if runtime.GOOS = "windows" {
		distPath += ".exe"
	}
	script.Exec(`
		go build  -ldflags
          -X 'spotexample/pkg/version.appName=${appname}'
		  -X 'spotexample/pkg/version.envPrefix=SPE_CALC'
		  -X 'spotexample/pkg/version.buildTag=${version}'
		  -X 'spotexample/pkg/version.buildHash=${commit}'
		  -X 'spotexample/pkg/version.buildBranch=${branch}'
		  -X 'spotexample/pkg/version.buildTime=$(date +"%Y-%m-%d %H:%M:%S %z")'
		  " -tags "mysql" -o ${name} ./cmd/spotexample-calc
	`).Stdout()
}

func must(str string, err error) string {
	if err != nil {
		log.Fatalln(err)
	}
	return str
}
