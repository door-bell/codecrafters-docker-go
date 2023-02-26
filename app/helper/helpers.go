package helper

import (
	"encoding/json"
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

func PrettyPrint(obj interface{}) string {
	formattedJson, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		DebugLog("Error formatting json", obj)
	}
	return string(formattedJson)
}
