package main

import (
	"encoding/json"
)

func JSON(data interface{}) string {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(jsonStr)
}
