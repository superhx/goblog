package blog

import (
	"fmt"
	"github.com/superhx/marker"
	"io"
	"os"
)

//RenderCategory ...
func RenderCategory(category []Article) {

	s := make([]Item, len(category))
	for i := range category {
		s[i] = category[i]
	}
	By(sortByDate).Sort(s)

	file, err := os.Create(outputDir + "template/category.htm")
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

//RenderArticle ...
func RenderArticle(mark *marker.MarkDown, output io.Writer) {
	marker.NewHTMLWriter(mark).WriteTo(output)
}

func sortByDate(i1, i2 Item) bool {
	a1, _ := i1.(Article)
	a2, _ := i2.(Article)
	return a1.Date.Before(a2.Date.Time)
}
