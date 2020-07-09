package app

import (
	"net/http"

	"github.com/cmd-ctrl-q/golang-microservices/mvc/controllers"
)

func StartApp() {
	// path and controller in charge of getting user info. ie handling path
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
