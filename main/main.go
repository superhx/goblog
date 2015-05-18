package main

import (
	"github.com/superhx/goblog"
)

func main() {
	var blog goblog.Blog
	blog.Transform()
	goblog.Server()
}
