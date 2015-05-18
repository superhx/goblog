package goblog

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

//Server ...
func Server() (err error) {
	err = http.ListenAndServe(":8080", http.FileServer(http.Dir(config.PublicDir)))
	if err != nil {
		log.Errorln("[Fail] fail to start server: ", err)
		return
	}
	log.Infoln("Server start")
	return
}
