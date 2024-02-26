package helper

import (
	"log"
)

func HelperError(err error, msg string) {
	if err != nil {
		log.Printf("\nerror: %v\nmessage: %s", err, msg)
		panic(err)
	}
}
