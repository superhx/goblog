package blog

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/superhx/mark"
)

//Generate ...
func Generate() (err error) {
	var blog Blog
	err = blog.Generate()
	if err != nil {
		log.Errorln(err)
	}
	return
}

//New ...
func New(title string, tags []string, content string, category bool) {
	article := make(map[string]interface{})
	article["title"] = title
	article["tags"] = tags
	article["date"] = time.Now().Format("2006/01/02|15:04:05")
	article["category"] = category
	bytes, _ := json.Marshal(article)
	str := "```\n" + string(bytes) + "\n```\n\n"
	path := fmt.Sprintf("%s/articles/%s.md", config.SourceDir, strings.Replace(title, "?", "", -1))
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Errorln(err)
		return
	}
	file.WriteString(str)
	file.WriteString(content)
	log.Infoln("Create a new blog: ", title)
}

//Blog ias ...
type Blog struct {
	articles []interface{}
	wg       sync.WaitGroup
}

//Generate ...
func (blog *Blog) Generate() (err error) {
	//generate static article html
	files, err := blog.files()

	defer func() {
		//generate category.json
		b, _ := json.Marshal(blog.articles)
		err = ioutil.WriteFile("category.json", b, os.ModePerm)
		if err != nil {
			log.Warnln("[Generate Fail]: category.json")
		}

		//generate static category html
		By(func(i1, i2 interface{}) bool {
			return i1.(*Article).Date.Time.After(i2.(*Article).Date.Time)
		}).Sort(blog.articles)

		renderCategory(blog.articles)

		//generate home page
		if len(blog.articles) != 0 {
			renderHomePage(getOutputPath(*(blog.articles[0].(*Article))))
		}
	}()

	if err != nil || len(files) == 0 {
		return
	}
	for _, file := range files {
		blog.wg.Add(1)
		go blog.generate(file)
	}

	blog.wg.Wait()

	return
}

func (blog *Blog) files() (files []os.FileInfo, err error) {

	old, err := ioutil.ReadDir(config.SourceDir + "/articles")
	if err != nil {
		log.Warnln("Can not open source dir: ", config.SourceDir)
		return
	}

	category, err := ioutil.ReadFile("category.json")
	var aticles []*Article

	//not init before or category.json broken
	if err != nil || json.Unmarshal(category, &aticles) != nil {
		log.Info("Generate all")
		os.MkdirAll(config.PublicDir+"/template", os.ModePerm)
		os.MkdirAll(config.PublicDir+"/css", os.ModePerm)
		os.MkdirAll(config.PublicDir+"/js", os.ModePerm)
		os.MkdirAll(config.PublicDir+"/img", os.ModePerm)
		for _, file := range old {
			if !file.IsDir() && file.Name()[0] != '.' {
				files = append(files, file)
			}
		}
		exec.Command("cp", "-R", config.SourceDir+"/data/", config.PublicDir).Run()
		err = nil
		return
	}

	m := make(map[string]*Article)
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
			blog.articles = append(blog.articles, m[name])
		} else {
			fmt.Println(article.ModTime(), " ", file.ModTime().UTC())
			os.Remove(config.PublicDir + "/" + getOutputPath(*article))
			files = append(files, file)
		}
		delete(m, name)
	}

	for name := range m {
		path := config.PublicDir + "/" + getOutputPath(*m[name])
		os.Remove(path)
		log.Infoln("[Remove]: ", path)
	}
	return
}

func (blog *Blog) generate(fileInfo os.FileInfo) {

	defer blog.wg.Done()

	//markdown input
	input, err := ioutil.ReadFile(config.SourceDir + "/articles/" + fileInfo.Name())
	if err != nil {
		log.Warnln("Can not open file: ", fileInfo.Name())
		return
	}

	//parse markdown to *Markdown obj
	markdown := mark.Mark(input)

	//extract article info
	article, err := getArticle(markdown, fileInfo)
	if err != nil {
		log.Error("[Format Error]: ", fileInfo.Name(), "; ", err)
		return
	}
	blog.articles = append(blog.articles, &article)

	markdown.Parts = markdown.Parts[1:]

	//transform markdown to html and output
	renderArticle(markdown, article)
}

func getArticle(markdown *mark.MarkDown, fileInfo os.FileInfo) (article Article, err error) {
	setting, ok := markdown.Parts[0].(*mark.Code)
	if !ok {
		err = errors.New("format error")
		return
	}
	err = json.Unmarshal([]byte(html.UnescapeString(setting.Text)), &article)
	if err != nil {
		log.Warningln(err)
		return
	}
	article.JSONFileInfo = &JSONFileInfo{fileInfo.Name(), fileInfo.Size(), fileInfo.Mode(), &JSONTime{fileInfo.ModTime()}, fileInfo.IsDir()}
	return
}

func getOutputPath(article Article) (outputPath string) {
	pubDate := article.Date
	fileName := article.Name()
	outputPath = strconv.Itoa(pubDate.Year()) + "/" + strconv.Itoa(int(pubDate.Month())) + "/" + strconv.Itoa(pubDate.Day()) + "/" + fileName[:len(fileName)-len(filepath.Ext(fileName))] + "/index.html"
	return
}
