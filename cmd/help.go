package main

import (
	"fmt"
)

var helpMap map[string]string

func init() {
	helpMap = map[string]string{
		"unkown": "Usage: ./goblog <command> \n\nCommands:\n generate|g\tGenerates static files\n server|s\tStart server to preview your website\ninit|i\tIntial workspace in current directory\nnew|n\tNew a article\n" +
			"For more help, you can use `./goblog help [command]` for the detailed information",
		"server":   "Usage: ./goblog server\nDescription: Start the server",
		"generate": "Usage: ./goblog generate\nDescription: Generate static files",
		"new":      "Usage: ./goblog new <title>\nDescription: Create a new article",
	}
}

func help(cmd string) {
	if str, exist := helpMap[cmd]; exist {
		fmt.Println(str)
	}
}
