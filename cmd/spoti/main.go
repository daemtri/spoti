package main

import (
	"github.com/duanqy/spoti/api"
	"github.com/duanqy/spoti/pkg/apispot"
)

func main() {
	apispot.Parser(api.Specific)
}
