package main

import (
	"bitbucket.org/psealock/rv/rvService"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

func main() {
	rvService.Console("Listening on port:8888")
	restful.Add(rvService.New())
	log.Fatal(http.ListenAndServe(":8888", nil))
}
