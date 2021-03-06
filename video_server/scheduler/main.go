package main

import (
	"net/http"

	"video_server/scheduler/taskrunner"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDleRecHandler)

	return router
}

func main() {
	go taskrunner.Start()

	router := RegisterHandlers()

	http.ListenAndServe(":9001", router)
}
