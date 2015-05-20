package goblog

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
)

//Server ...
func Server(port int) (err error) {
	err = http.ListenAndServe(":"+strconv.Itoa(port), http.FileServer(http.Dir(config.PublicDir)))
	if err != nil {
		log.Errorln("[Fail] fail to start server: ", err)
		return
	}
	log.Infoln("Server start")
	return
}
