package version

import (
	"fmt"
	"runtime"
)

var (
	appName     string
	buildTag    string
	buildHash   string
	buildBranch string
	buildTime   string
	envPrefix   string
)

func AppName() string {
	return appName
}
func BuildTag() string {
	return buildTag
}

func BuildHash() string {
	return buildHash
}

func BuildBranch() string {
	return buildBranch
}

func BuildTime() string {
	return buildTime
}

func EnvPrefix() string {
	return envPrefix
}

func Print() {
	fmt.Printf("%s version %s hash=%s branch=%s time=%s (go version %s)\n",
		appName,
		buildTag,
		buildHash,
		buildBranch,
		buildTime,
		runtime.Version(),
	)
}
