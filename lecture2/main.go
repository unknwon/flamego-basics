package main

import (
	"log"
	"net/http"

	"github.com/flamego/flamego"
)

func main() {
	f := flamego.Classic()
	f.Get("/{**}", printRequestPath)

	log.Println("Server is running...")
	log.Println(http.ListenAndServe("0.0.0.0:2830", f))
}

func printRequestPath(r *http.Request) string {
	return "The request path is: " + r.RequestURI
}
