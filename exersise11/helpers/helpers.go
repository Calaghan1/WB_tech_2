package helpers

import (
	"log"
)
func CheckError(err error) bool {
	if err != nil {
		log.Println(err.Error())
		return false
	} else {
		return true
	}
}