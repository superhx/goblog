package goblog

import (
	"fmt"
	"github.com/superhx/mark"
	"io"
	"os"
)

//RenderCategory ...
func renderCategory(category []Article) {

	s := make([]interface{}, len(category))
	for i := range category {
		s[i] = category[i]
	}
	By(sortByDate).Sort(s)

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

//RenderArticle ...
func renderArticle(node mark.Node, article Article, output io.Writer) {
	mark.NewHTMLWriter(node).WriteTo(output)
}

func sortByDate(i1, i2 interface{}) bool {
	a1, _ := i1.(Article)
	a2, _ := i2.(Article)
	return a1.Date.Before(a2.Date.Time)
}
