package errorhandler

import "log"

func HanderError(m string, err error) {
	if err != nil {
		log.Fatal(err)
	}
}
