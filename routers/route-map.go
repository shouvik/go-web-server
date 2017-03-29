package routers

import (
	"go-web-server/controllers"

	"github.com/codegangsta/negroni"
)

type routeMap struct {
	routePath          string
	controllerFunction negroni.HandlerFunc
	requestType        string
}

var routeMapObject []routeMap

func initRouteMaps() {
	routeMapObject = append(routeMapObject, routeMap{"/sample", controllers.SampleController, "GET"})
}
