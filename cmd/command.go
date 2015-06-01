package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/kardianos/osext"
	"github.com/skratchdot/open-golang/open"
	"github.com/superhx/goblog"
	"github.com/zenazn/goji/graceful"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

var config goblog.Config

func init() {
	config = goblog.GetConfig()
}

func main() {
	if len(os.Args) == 1 {
		help("unkown")
		return
	}

	switch cmd := os.Args[1]; {
	case cmd == "server" || cmd == "s":
		server()
		defer graceful.Wait()
	case cmd == "generate" || cmd == "g":
		generate()
	case cmd == "init" || cmd == "i":
		workspace()
	case cmd == "new" || cmd == "n":
		article(os.Args[2:])
	case cmd == "help" || cmd == "h":
		if len(os.Args) > 2 {
			help(os.Args[2])
		}
	default:
		help("unkown")
	}
}

func server() {
	go goblog.Server(8001)
	open.Run("http://localhost:8001/new.html")
}

func generate() {
	goblog.Generate()
	exec.Command("cp", "-R", config.SourceDir+"/data/", config.PublicDir).Run()
	dir, _ := osext.ExecutableFolder()
	exec.Command("cp", "-R", dir+"/theme/data/", config.PublicDir).Run()
}

func workspace() {
	config := goblog.GetConfig()
	jconfig, _ := json.MarshalIndent(config, "", "    ")
	ioutil.WriteFile("config.json", jconfig, os.ModePerm)
	os.MkdirAll(config.PublicDir, os.ModePerm)
	os.MkdirAll(config.SourceDir+"/articles", os.ModePerm)
	os.MkdirAll(config.SourceDir+"/data", os.ModePerm)
	fmt.Println("Init workspace done")
}

func article(args []string) {
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
