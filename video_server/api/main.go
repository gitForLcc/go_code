package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := &middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	validateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func RegisterHandles() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}

// handler -> validation{1.request 2.user}->business logic->reponse
// main->middleware->defs(message, err)->handlers->dbops->response
func main() {
	router := RegisterHandles()
	mh := NewMiddleWareHandler(router)

	http.ListenAndServe(":8000", mh)
}
