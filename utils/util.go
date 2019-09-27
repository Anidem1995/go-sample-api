package utils

import (
	"encoding/json"
	_ "encoding/json"
	"net/http"
	_ "net/http"
)

func Message(status bool, message string) map[string] interface{} {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(w http.ResponseWriter, data map[string] interface{}) {
	w.Header().Add("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(data)
}