package main

import (
	"go-web-server/routers"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
)

func main() {
	// settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	log.Println("Starting Go Server on Port 5000")
	http.ListenAndServe(":5000", n)
}
