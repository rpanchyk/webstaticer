package utils

import (
	"encoding/json"
	"log"
)

func ToPrettyString(obj interface{}) string {
	pretty, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Println(err)
		return "error while prettifying"
	}
	return string(pretty)
}
