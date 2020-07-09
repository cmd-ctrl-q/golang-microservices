package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cmd-ctrl-q/golang-microservices/mvc/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	// userID := req.URL.Query().Get("user_id")
	// log.Printf("About to process user_id %v", userID)

	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		// resp.WriteHeader(http.StatusNotFound)
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("user_id must be a number"))
		return
	}

	// process user id
	user, err := services.GetUser(userID)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		// handle the error and erturn to the client
		return
	}

	// else, return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
