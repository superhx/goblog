package blog

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/zenazn/goji/graceful"
	"net/http"
	"os/exec"
	"strconv"
)

//Server ...
func Server(port int) {
	log.Info("Server start. Please visit http://localhost:" + strconv.Itoa(port))
	log.Infoln("Press ctrl-c to stop")
	http.Handle("/dashboard/blog/new", http.HandlerFunc(newBlog))
	http.Handle("/dashboard/generate", http.HandlerFunc(generate))
	http.Handle("/dashboard/deploy", http.HandlerFunc(deploy))
	fmt.Println(ThemeDir + "/dashboard")
	http.Handle("/dashboard/", http.StripPrefix("/dashboard/", http.FileServer(http.Dir(ThemeDir+"/dashboard"))))
	http.Handle("/", http.FileServer(http.Dir(config.PublicDir)))
	if err := graceful.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		log.Errorln("[Fail] fail to start server: ", err)
	}
}

func newBlog(w http.ResponseWriter, req *http.Request) {
	title := req.PostFormValue("title")
	tags := req.Form["tags[]"]
	fmt.Println(tags)
	content := req.PostFormValue("content")
	category, _ := strconv.ParseBool(req.PostFormValue("category"))
	New(title, tags, content, category)
	Generate()
	res, _ := json.Marshal(map[string]string{"status": "success"})
	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}

func generate(w http.ResponseWriter, req *http.Request) {
	Generate()
	res, _ := json.Marshal(map[string]string{"status": "success"})
	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}

func deploy(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	err := exec.Command("git", "push").Run()
	var res []byte
	if err != nil {
		log.Errorln(err)
		res, _ = json.Marshal(map[string]string{"status": "fail", "error": err.Error()})
	} else {
		res, _ = json.Marshal(map[string]string{"status": "success"})
	}
	w.Write(res)
}
