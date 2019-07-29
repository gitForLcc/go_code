package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(b *ApiBody, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var resp *http.Response
	var err error

	u := b.Url + ps.ByName("proxyPath")
	log.Println(r.URL, "=>", u)
	req, _ := http.NewRequest(r.Method, u, r.Body)
	req.Header = r.Header
	resp, err = httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	normalResponse(w, resp)

}

func normalResponse(w http.ResponseWriter, r *http.Response) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		re, _ := json.Marshal(ErrorInternalFaults)
		w.WriteHeader(500)
		io.WriteString(w, string(re))
		return
	}

	w.WriteHeader(r.StatusCode)
	io.WriteString(w, string(res))

}
