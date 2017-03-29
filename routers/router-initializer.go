package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func mapRoutesWithControllers(router *mux.Router) *mux.Router {
	for _, mapObject := range routeMapObject {
		router.Handle(mapObject.routePath,
			negroni.New(
				mapObject.controllerFunction,
			)).Methods(mapObject.requestType)
	}
	return router
}
