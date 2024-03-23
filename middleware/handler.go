package middleware

import "log"

func CheckError(err error) {
	if err != nil {
		log.Printf(err.Error())
	}
}
