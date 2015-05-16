package blog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var config Config

//Config ...
type Config struct {
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	Author    string `json:"author"`
	Email     string `json:"email"`
	URL       string `json:"url"`
	Root      string `json:"root"`
	SourceDir string `json:"sourceDir"`
	PublicDir string `json:"publicDir"`
	Disqus    string `json:"disqus"`
	Analyze   string `json:"analyze"`
}

func init() {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("fail to read config.json")
	}
	json.Unmarshal(bytes, &config)
}
