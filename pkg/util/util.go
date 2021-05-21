package util

import (
	"github.com/rs/zerolog/log"
)

// HandleErr is a generic error handler
func HandleErr(err error, message string) {
	if err != nil {
		log.Fatal().Err(err).Msg(message)
	}
}
