package goblog

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/kardianos/osext"
	"github.com/superhx/mark"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)

var templateDir string
var homeTmpl *template.Template
var blogTmpl *template.Template

func init() {
	dir, _ := osext.ExecutableFolder()
	templateDir = dir + "/../theme/template"

	homeTmpl, _ = template.ParseFiles(templateDir + "/home.htm")
	blogTmpl, _ = template.ParseFiles(templateDir + "/article.htm")
}

func renderCategory(category []interface{}) {
	file, err := os.Create(config.PublicDir + "/template/category.htm")
	defer file.Close()
	if err != nil {
		fmt.Println("Create Categoty template fail")
	}
	file.WriteString("<nav><ul>")

	data := make([]map[string]interface{}, 0, len(category))
	for _, item := range category {
		article := item.(*Article)
		link := config.Root + "/" + getOutputPath(*article)
		file.WriteString("<li><a href=\"" + link + "\">" + article.Title + "</a></li>")

		data = append(data, map[string]interface{}{"title": article.Title, "tags": article.Tags, "date": article.Date, "link": getOutputPath(*article)})
	}
	file.WriteString("</ul></nav>")

	json, _ := json.Marshal(data)
	err = ioutil.WriteFile(config.PublicDir+"/category.json", json, os.ModePerm)
	if err != nil {
		log.Errorln(err)
	}
}

func renderArticle(node mark.Node, article Article) {
	outputPath := config.PublicDir + "/" + getOutputPath(article)
	os.MkdirAll(path.Dir(outputPath), os.ModePerm)
	output, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer output.Close()
	if err != nil {
		log.Errorln(err)
		return
	}
	buf := bytes.NewBufferString("")
	mark.NewHTMLWriter(node).WriteTo(buf)
	blogTmpl.Execute(output, map[string]interface{}{"Info": article, "Content": buf})
	if err != nil {
		log.Errorln(err)
		return
	}
	log.Infoln("[Generate]: ", outputPath)
}

func renderHomePage(first string) {
	output, err := os.Create(config.PublicDir + "/index.html")
	defer output.Close()
	if err != nil {
		log.Errorln(err)
		return
	}
	homeTmpl.Execute(output, map[string]string{"First": first})
}
