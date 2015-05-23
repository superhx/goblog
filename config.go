package goblog

import (
	"encoding/json"
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
		config.SourceDir = "source"
		config.PublicDir = "public"
		return
	}
	json.Unmarshal(bytes, &config)
}

//GetConfig ...
func GetConfig() Config {
	return config
}
