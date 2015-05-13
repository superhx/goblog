package blog

import (
	"encoding/json"
	"fmt"
	"github.com/superhx/marker"
	"html"
	"io/ioutil"
	"os"
)

//Blog is ...
type Blog struct {
	articles  []Article
	parallels chan bool
}

//Transform ...
func (blog *Blog) Transform() {
	files, err := ioutil.ReadDir(inputPath)
	if err != nil {
		fmt.Println("Can not read dir:" + inputPath)
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
}

func (blog *Blog) transform(fileInfo os.FileInfo) {
	input, err := ioutil.ReadFile(inputPath + fileInfo.Name())
	if err != nil {
		fmt.Println("Can not read file:" + fileInfo.Name())
		return
	}
	mark := marker.Mark(input)

	article := getArticle(mark)
	mark.Parts[0] = &marker.Heading{Depth: 1, Text: &marker.Text{Parts: []marker.Node{&marker.InlineText{Text: article.Title}}}}
	fmt.Println(article.Date)
	<-blog.parallels
}

func getArticle(mark *marker.MarkDown) (article Article) {
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
