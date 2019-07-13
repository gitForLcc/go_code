package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandles() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}

// handler -> validation{1.request 2.user}->business logic->reponse

func main() {
	router := RegisterHandles()

	http.ListenAndServe(":8000", router)
}
