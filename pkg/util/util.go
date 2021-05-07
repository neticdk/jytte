package util

import (
	"log"
)

// HandleErr is a generic error handler
func HandleErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
