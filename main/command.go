package main

import (
	"fmt"
	"github.com/superhx/goblog"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		unkown()
		return
	}
	//	fmt.Println(os.Getwd())
	switch cmd := os.Args[1]; {
	case cmd == "server" || cmd == "s":
		goblog.Server(8001)
	case cmd == "generate" || cmd == "g":
		goblog.Generate()
		fmt.Println("generate")
	default:
		unkown()
	}
}

func unkown() {
	fmt.Println("Usage: ./goblog <command> \n\n" + "Commands:\n generate|g\tGenerates static files\n" +
		" server|s\tStart server to preview your website")
}
