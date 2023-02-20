package helper

import (
	"log"
	"os"
)

func IsDebug() bool {
	return os.Getenv("DEBUG") != ""
}

func DebugLog(v ...interface{}) {
	if IsDebug() {
		log.Println(v...)
	}
}
