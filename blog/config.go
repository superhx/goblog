package blog

import (
	"encoding/json"
	"io/ioutil"
)

var config Config

func init() {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		config.SourceDir = "source"
		config.PublicDir = "public"
		return
	}
	json.Unmarshal(bytes, &config)
}

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

//GetConfig ...
func GetConfig() Config {
	return config
}
