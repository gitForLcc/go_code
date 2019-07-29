package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", homeHandler)

	router.POST("/api", apiHandler)

	router.POST("upload/vid-id", proxyHandler)

	return router
}

func main() {
	router := RegisterHandlers()
	http.ListenAndServe(":8000", router)
}
