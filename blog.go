package blog

import (
	"encoding/json"
	"fmt"
	"github.com/superhx/marker"
	"html"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

const (
	articleParallelCount = 100
)

//Blog is ...
type Blog struct {
	Articles  []*Article
	parallels chan bool
}

//Transform ...
func (blog *Blog) Transform() {
	os.RemoveAll(outputDir)
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
	fmt.Println(string(b))
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
	article := getArticle(mark, fileInfo)
	blog.Articles = append(blog.Articles, &article)
	// d, _ := article.Date.MarshalJSON()
	// fmt.Println(string(d))

	//set markdown title
	mark.Parts[0] = &marker.Heading{Depth: 1, Text: &marker.Text{Parts: []marker.Node{&marker.InlineText{Text: article.Title}}}}

	//create output path and output file
	pubDate := article.Date
	fileName := fileInfo.Name()
	outputPath := outputDir + strconv.Itoa(pubDate.Year()) + "/" + strconv.Itoa(int(pubDate.Month())) + "/" + strconv.Itoa(pubDate.Day()) + "/" + fileName[:len(fileName)-len(filepath.Ext(fileName))]
	os.MkdirAll(outputPath, os.ModePerm)
	output, err := os.OpenFile(outputPath+"/index.html", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer func() { output.Close() }()
	if err != nil {
		fmt.Println("Can not create file: " + outputPath)
	}

	//transform markdown to html and output
	marker.NewHTMLWriter(mark).WriteTo(output)
}

func getArticle(mark *marker.MarkDown, fileInfo os.FileInfo) (article Article) {
	setting, ok := mark.Parts[0].(*marker.Code)
	if !ok {
		fmt.Println("Format error")
		return
	}
	json.Unmarshal([]byte(html.UnescapeString(setting.Text)), &article)
	return
}

func getOutputPath(article Article) string {
	return ""
}
