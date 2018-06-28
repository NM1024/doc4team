package tools

import (
	"log"
)

// Err Log
func LogErr(err interface{}) {
	if err != nil {
		// log.Fatal(err)
		log.Println(err)
	}
}
