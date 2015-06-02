package goblog

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	graceful "github.com/zenazn/goji/graceful"
	"net/http"
	"strconv"
)

//Server ...
func Server(port int) {
	log.Info("Server start. Please visit http://localhost:" + strconv.Itoa(port))
	log.Infoln("Press ctrl-c to stop")
	http.Handle("/dashboard/blog/new", http.HandlerFunc(newBlog))
	http.Handle("/", http.FileServer(http.Dir(config.PublicDir)))
	if err := graceful.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		log.Errorln("[Fail] fail to start server: ", err)
	}
}

func newBlog(w http.ResponseWriter, req *http.Request) {
	title := req.PostFormValue("title")
	tags := req.Form["tags[]"]
	content := req.PostFormValue("content")
	New(title, tags, content)
	res, _ := json.Marshal(map[string]string{"status": "success"})
	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}
