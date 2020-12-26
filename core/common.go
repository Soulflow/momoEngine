package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MessageBody struct {
	Total   int         `json:"total"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"msg"`
}

func (mb *MessageBody) ResponseCommon(w http.ResponseWriter, statusCode int) {
	jsonStr, err := json.Marshal(mb)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	w.WriteHeader(statusCode)
	_, _ = fmt.Fprintf(w, string(jsonStr))
}
