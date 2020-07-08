package controllers

import (
	"net/http"

	"github.com/cmd-ctrl-q/golang-microservices/mvc/controllers"
)

func GetUser() {
	http.HandleFunc("/users", controllers.GetUser)
}
