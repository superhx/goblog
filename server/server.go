package server

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/superhx/goblog/blog"
	"github.com/zenazn/goji/graceful"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

//Server ...
func Server(port int) {
	log.Info("Server start. Please visit http://localhost:" + strconv.Itoa(port))
	log.Infoln("Press ctrl-c to stop")
	http.Handle("/dashboard/blog/new", http.HandlerFunc(newBlogHandle))
	http.Handle("/dashboard/generate", http.HandlerFunc(generateHandle))
	http.Handle("/dashboard/deploy", http.HandlerFunc(deployHandle))
	http.Handle("/dashboard/", http.StripPrefix("/dashboard/", http.FileServer(http.Dir(blog.ThemeDir+"/dashboard"))))
	var config = blog.GetConfig()
	http.Handle("/", http.FileServer(http.Dir(config.PublicDir)))
	fmt.Println(blog.ThemeDir + "/dashboard")
	if err := graceful.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		log.Errorln("[Fail] fail to start server: ", err)
	}
}

func newBlogHandle(w http.ResponseWriter, req *http.Request) {
	title := req.PostFormValue("title")
	tags := req.Form["tags[]"]
	fmt.Println(tags)
	content := req.PostFormValue("content")
	category, _ := strconv.ParseBool(req.PostFormValue("category"))
	blog.New(title, tags, content, category)
	blog.Generate()
	res, _ := json.Marshal(map[string]string{"status": "success"})
	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}

func generateHandle(w http.ResponseWriter, req *http.Request) {
	blog.Generate()
	res, _ := json.Marshal(map[string]string{"status": "success"})
	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}

func deployHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	err := Deploy()
	var res []byte
	if err != nil {
		log.Errorln(err)
		res, _ = json.Marshal(map[string]string{"status": "fail", "error": err.Error()})
	} else {
		res, _ = json.Marshal(map[string]string{"status": "success"})
	}
	w.Write(res)
}

// Deploy ...
func Deploy() (err error) {
	os.Chdir("public")
	defer func() {
		os.Chdir("..")
	}()
	if err = command("git", "add", "*"); err != nil {
		return
	}
	if err = command("git", "commit", "-a", "-m", "~"); err != nil {
		return
	}
	if err = command("git", "push", "origin", "master"); err != nil {
		return
	}
	return
}

func command(name string, args ...string) (err error) {
	out, err := exec.Command(name, args...).Output()
	if err != nil {
		log.Errorln(err, ":", string(out))
		return
	}
	log.Infoln(string(out))
	return
}
