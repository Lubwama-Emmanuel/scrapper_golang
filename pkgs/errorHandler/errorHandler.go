package errorhandler

import log "github.com/sirupsen/logrus"

func HanderError(m string, err error) {
	if err != nil {
		log.Fatal(err)
	}
}
