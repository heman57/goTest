package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	APIRoot := func(w http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		io.WriteString(w, req.URL.Path)
		io.WriteString(w, string(body))
	}
	http.HandleFunc("/", APIRoot)
	http.ListenAndServe(":80", nil)
}
