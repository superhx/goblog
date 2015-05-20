package main

import (
	"github.com/superhx/goblog"
)

func main() {
	goblog.Generate()
	goblog.Server(8080)
}
