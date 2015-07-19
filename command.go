package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"
	"github.com/superhx/goblog/blog"
	"github.com/superhx/goblog/server"
	"github.com/zenazn/goji/graceful"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

var config blog.Config

func init() {
	config = blog.GetConfig()
}

func main() {
	start := time.Now()
	defer func() {
		end := time.Now()
		fmt.Printf("The call took %v to run.\n", end.Sub(start))
	}()
	if len(os.Args) == 1 {
		help("unkown")
		return
	}
	switch cmd := os.Args[1]; {
	case cmd == "server" || cmd == "s":
		startServer()
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
	case cmd == "deploy" || cmd == "d":
		server.Deploy()
	default:
		help("unkown")
	}
}

func startServer() {
	go server.Server(8001)
	open.Run("http://localhost:8001/dashboard")
}

func generate() {
	blog.Generate()
}

func workspace() {
	config := blog.GetConfig()
	jconfig, _ := json.MarshalIndent(config, "", "    ")
	ioutil.WriteFile("config.json", jconfig, os.ModePerm)
	os.MkdirAll(config.PublicDir, os.ModePerm)
	os.MkdirAll(config.SourceDir+"/articles", os.ModePerm)
	os.MkdirAll(config.SourceDir+"/data", os.ModePerm)
	err := exec.Command("cp", "-R", blog.ThemeDir+"/data/", config.PublicDir).Run()
	if err != nil {
		log.Errorln(err)
		return
	}
	// err = exec.Command("cp", "-R", blog.ThemeDir+"/dashboard", config.PublicDir).Run()
	// if err != nil {
	// 	log.Errorln(err)
	// 	return
	// }
	log.Infoln("Init workspace done")
}

func article(args []string) {
	if len(args) == 0 {
		help("new")
		return
	}
	title := args[0]
	blog.New(title, []string{}, "", false)
}
