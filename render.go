package goblog

import (
	"bytes"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/kardianos/osext"
	"github.com/superhx/mark"
	"os"
	"path"
	"text/template"
)

var templateDir string
var homeTmpl *template.Template
var blogTmpl *template.Template

func init() {
	dir, _ := osext.ExecutableFolder()
	templateDir = dir + "../theme/template"

	homeTmpl, _ = template.ParseFiles(templateDir + "/home.htm")
	blogTmpl, _ = template.ParseFiles(templateDir + "/article.htm")
}

func renderCategory(category []Article) {
	file, err := os.Create(config.PublicDir + "/template/category.htm")
	defer file.Close()
	if err != nil {
		fmt.Println("Create Categoty template fail")
	}
	file.WriteString("<nav><ul>")
	for _, item := range category {
		link := config.Root + "/" + getOutputPath(item)
		file.WriteString("<li><a href=\"" + link + "\">" + item.Title + "</a></li>")
	}
	file.WriteString("</ul></nav>")
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

func sortArticle(category []Article, by By) {
	s := make([]interface{}, len(category))
	for i := range category {
		s[i] = category[i]
	}
	By(by).Sort(s)
}

func byDateR(i1, i2 interface{}) bool {
	a1, _ := i1.(Article)
	a2, _ := i2.(Article)
	return a1.Date.After(a2.Date.Time)
}
