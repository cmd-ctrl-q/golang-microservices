package controllers

import (
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	// userID := req.URL.Query().Get("user_id")
	// log.Printf("About to process user_id %v", userID)

	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {

	}
}
