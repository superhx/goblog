package blog

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/superhx/marker"
	"html"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

var _ = flag.ContinueOnError
var _ = glog.CopyStandardLogTo

const (
	articleParallelCount = 100
)

//Blog is ...
type Blog struct {
	Articles  []Article
	parallels chan bool
}

//Transform ...
func (blog *Blog) Transform() {

	files := blog.files()

	for _, file := range files {
		fmt.Println("Parse file:" + file.Name())
		blog.parallels <- true
		go blog.transform(file)
	}

	for i := 0; i < articleParallelCount; i++ {
		blog.parallels <- false
	}

	b, _ := json.Marshal(blog.Articles)
	err := ioutil.WriteFile(config.PublicDir+"/category.json", b, os.ModePerm)
	if err != nil {
		fmt.Println("Write category fail!")
	}

	RenderCategory(blog.Articles)

}

func (blog *Blog) transform(fileInfo os.FileInfo) {

	defer func() { <-blog.parallels }()

	//markdown input
	input, err := ioutil.ReadFile(config.SourceDir + "/" + fileInfo.Name())
	if err != nil {
		fmt.Println("Can not read file: " + fileInfo.Name())
		return
	}

	//parse markdown to *Markdown obj
	mark := marker.Mark(input)

	//extract article info
	article := GetArticle(mark, fileInfo)
	blog.Articles = append(blog.Articles, article)

	//set markdown title
	mark.Parts[0] = &marker.Heading{Depth: 1, Text: &marker.Text{Parts: []marker.Node{&marker.InlineText{Text: article.Title}}}}

	//create output dir and output file
	outputPath := GetOutputPath(article)
	os.MkdirAll(path.Dir(outputPath), os.ModePerm)
	output, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer output.Close()
	if err != nil {
		fmt.Println("Can not create file: " + outputPath)
	}

	//transform markdown to html and output
	RenderArticle(mark, output)
}

func (blog *Blog) files() (files []os.FileInfo) {

	old, err := ioutil.ReadDir(config.SourceDir)
	if err != nil {
		fmt.Println("Can not read dir:" + config.SourceDir)
		return
	}

	category, err := ioutil.ReadFile(config.PublicDir + "/category.json")
	var aticles []Article

	//not init before or category.json broken
	if err != nil || json.Unmarshal(category, &aticles) != nil {
		fmt.Println("Init from empty")
		os.RemoveAll(config.PublicDir)
		os.MkdirAll(config.PublicDir+"/template", os.ModePerm)
		os.MkdirAll(config.PublicDir+"/css", os.ModePerm)
		os.MkdirAll(config.PublicDir+"/js", os.ModePerm)
		os.MkdirAll(config.PublicDir+"/img", os.ModePerm)
		for _, file := range old {
			if !file.IsDir() && file.Name()[0] != '.' {
				files = append(files, file)
			}
		}
		return files
	}

	m := make(map[string]Article)
	for _, a := range aticles {
		m[a.Name()] = a
	}

	for _, file := range old {
		name := file.Name()
		if file.IsDir() || name[0] == '.' {
			continue
		}

		article, exist := m[name]
		if !exist {
			files = append(files, file)
			continue
		}

		if article.ModTime().Equal(file.ModTime()) {
			blog.Articles = append(blog.Articles, m[name])
		} else {
			fmt.Println(article.ModTime(), " ", file.ModTime().UTC())
			os.Remove(GetOutputPath(article))
			files = append(files, file)
		}
		delete(m, name)
	}

	for name := range m {
		fmt.Println("remove file")
		os.Remove(GetOutputPath(m[name]))
	}

	return
}

//GetArticle ...
func GetArticle(mark *marker.MarkDown, fileInfo os.FileInfo) (article Article) {
	setting, ok := mark.Parts[0].(*marker.Code)
	if !ok {
		fmt.Println("Format error")
		return
	}
	json.Unmarshal([]byte(html.UnescapeString(setting.Text)), &article)
	article.JSONFileInfo = &JSONFileInfo{fileInfo.Name(), fileInfo.Size(), fileInfo.Mode(), &JSONTime{fileInfo.ModTime()}, fileInfo.IsDir()}
	return
}

//GetOutputPath ...
func GetOutputPath(article Article) (outputPath string) {
	pubDate := article.Date
	fileName := article.Name()
	outputPath = config.PublicDir + "/" + strconv.Itoa(pubDate.Year()) + "/" + strconv.Itoa(int(pubDate.Month())) + "/" + strconv.Itoa(pubDate.Day()) + "/" + fileName[:len(fileName)-len(filepath.Ext(fileName))] + "/index.html"
	return
}
