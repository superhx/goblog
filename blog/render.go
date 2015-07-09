package blog

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

//ThemeDir
var ThemeDir string

var templateDir string
var homeTmpl *template.Template
var blogTmpl *template.Template

var plus = template.FuncMap{
	"plus1": func(x int) int {
		return x + 1
	},
}

func init() {
	ThemeDir, _ = osext.ExecutableFolder()
	ThemeDir += "/../src/github.com/superhx/goblog/theme"
	templateDir = ThemeDir + "/template"
	homeTmpl, _ = template.ParseFiles(templateDir + "/home.htm")
	blogTmpl, _ = template.ParseFiles(templateDir + "/article.htm")
	// // blogTmpl, err = template.ParseFiles(templateDir + "/article.htm")
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }
	// // blogTmpl = blogTmpl.Funcs(plus)
	blogTmpl = template.Must(new(template.Template).Funcs(plus).ParseFiles(templateDir + "/article.htm"))
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
	output, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	defer output.Close()
	if err != nil {
		log.Errorln(err)
		return
	}
	buf := bytes.NewBufferString("")
	mark.NewHTMLWriter(node).WriteTo(buf)
	blogTmpl.ExecuteTemplate(output, "article.htm", map[string]interface{}{"Info": article, "Content": buf, "Disqus": config.Disqus})
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
