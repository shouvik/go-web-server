package routers

import "github.com/gorilla/mux"

// InitRoutes : this function initializes the routes mentioned in the routes files
func InitRoutes() *mux.Router {
	initRouteMaps()
	router := mux.NewRouter()
	router = mapRoutesWithControllers(router)
	return router
}
