package errorhandler

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func HanderError(incomingErr error) {
	if incomingErr != nil {
		err := fmt.Errorf("error occurred: %w", incomingErr)
		log.Error(err)
	}
}
