package goblog

import (
	log "github.com/Sirupsen/logrus"
	graceful "github.com/zenazn/goji/graceful"
	"net/http"
	"strconv"
)

//Server ...
func Server(port int) {
	log.Info("Server start. Please visit http://localhost:" + strconv.Itoa(port))
	log.Infoln("Press ctrl-c to stop")
	if err := graceful.ListenAndServe(":"+strconv.Itoa(port), http.FileServer(http.Dir(config.PublicDir))); err != nil {
		log.Errorln("[Fail] fail to start server: ", err)
	}
}
