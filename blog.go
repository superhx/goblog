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
	context              = ""
)

//Blog is ...
type Blog struct {
	Articles  []Article
	parallels chan bool
}

//Transform ...
func (blog *Blog) Transform() {

	os.RemoveAll(outputDir)
	os.MkdirAll(outputDir+"template", os.ModePerm)
	os.MkdirAll(outputDir+"css", os.ModePerm)
	os.MkdirAll(outputDir+"js", os.ModePerm)
	os.MkdirAll(outputDir+"img", os.ModePerm)

	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		fmt.Println("Can not read dir:" + inputDir)
		return
	}

	for _, file := range files {
		if !file.IsDir() && file.Name()[0] != '.' {
			blog.parallels <- true
			go blog.transform(file)
		}
	}

	for i := 0; i < articleParallelCount; i++ {
		blog.parallels <- false
	}

	b, _ := json.Marshal(blog.Articles)
	err = ioutil.WriteFile(outputDir+"category.json", b, os.ModePerm)
	if err != nil {
		fmt.Println("Write category fail!")
	}

	GenerateCategory(blog.Articles)

}

func (blog *Blog) transform(fileInfo os.FileInfo) {

	defer func() { <-blog.parallels }()

	//markdown input
	input, err := ioutil.ReadFile(inputDir + fileInfo.Name())
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
	marker.NewHTMLWriter(mark).WriteTo(output)
}

//GetArticle ...
func GetArticle(mark *marker.MarkDown, fileInfo os.FileInfo) (article Article) {
	setting, ok := mark.Parts[0].(*marker.Code)
	if !ok {
		fmt.Println("Format error")
		return
	}
	json.Unmarshal([]byte(html.UnescapeString(setting.Text)), &article)
	article.Origin = &JSONFileInfo{fileInfo.Name(), fileInfo.Size(), fileInfo.Mode(), &JSONTime{fileInfo.ModTime()}, fileInfo.IsDir()}
	return
}

//GetOutputPath ...
func GetOutputPath(article Article) (outputPath string) {
	pubDate := article.Date
	fileName := article.Origin.Name
	outputPath = outputDir + strconv.Itoa(pubDate.Year()) + "/" + strconv.Itoa(int(pubDate.Month())) + "/" + strconv.Itoa(pubDate.Day()) + "/" + fileName[:len(fileName)-len(filepath.Ext(fileName))] + "/index.html"
	return
}

//GenerateCategory ...
func GenerateCategory(category []Article) {
	file, err := os.Create(outputDir + "template/category.htm")
	fmt.Println(outputDir + "template/category.htm")
	defer file.Close()
	if err != nil {
		fmt.Println("Create Categoty template fail")
	}
	file.WriteString("<nav><ul>")
	for _, item := range category {
		link := context + GetOutputPath(item)
		file.WriteString("<li><a href=\"" + link + "\">" + item.Title + "</a></li>")
	}
	file.WriteString("</ul></nav>")
}
