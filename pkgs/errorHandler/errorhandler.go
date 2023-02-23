package errorHandler

import "log"

func HanderError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
