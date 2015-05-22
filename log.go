package goblog

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}
