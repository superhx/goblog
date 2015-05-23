package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/superhx/goblog"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		help("unkown")
		return
	}
	//	fmt.Println(os.Getwd())
	switch cmd := os.Args[1]; {
	case cmd == "server" || cmd == "s":
		goblog.Server(8001)
	case cmd == "generate" || cmd == "g":
		goblog.Generate()
	case cmd == "init" || cmd == "i":
		initWorkspace()
	case cmd == "new" || cmd == "n":
		newArticle(os.Args[2:])
	case cmd == "help" || cmd == "h":
		if len(os.Args) > 2 {
			help(os.Args[2])
		}
	default:
		help("unkown")
	}
}

func initWorkspace() {
	config := goblog.GetConfig()
	jconfig, _ := json.MarshalIndent(config, "", "    ")
	ioutil.WriteFile("config.json", jconfig, os.ModePerm)
	os.MkdirAll(config.PublicDir, os.ModePerm)
	os.MkdirAll(config.SourceDir, os.ModePerm)
	fmt.Println("Init workspace done")
}

func newArticle(args []string) {
	if len(args) == 0 {
		help("new")
		return
	}
	config := goblog.GetConfig()
	title := args[0]
	article := fmt.Sprintf(`{"title":"%s", "date":"%s", "tag":%s}`, title, time.Now().Format("2006/01/02|15:04:05"), []string{})
	article = "```\n" + article + "\n```"
	path := fmt.Sprintf("%s/%s.md", config.SourceDir, title)
	if err := ioutil.WriteFile(path, []byte(article), os.ModePerm); err != nil {
		log.Errorf("Cannot create %s", path)
		return
	}

	log.Infof("File created at %s", path)
}
