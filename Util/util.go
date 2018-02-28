package Util

import "log"

func Must(err error) {
	if err != nil {
		log.Println("[Panic Server] Something went wrong")
	}
}