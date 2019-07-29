package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/*proxyPath", proxyHandler)

	return router
}

func main() {
	router := RegisterHandlers()
	http.ListenAndServe(":8001", router)
}
