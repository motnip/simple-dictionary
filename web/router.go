package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/motnip/sermo/system"
)

type Router struct {
	router *mux.Router
	log    *system.SermoLog
}

func NewRouter() *Router {
	return &Router{
		router: mux.NewRouter().StrictSlash(true),
		log:    system.NewLog(),
	}
}

func (r *Router) RouterStart() {
	r.log.LogInfo("Server started... ")
	log.Fatal(http.ListenAndServe(":3000", r.router))
}

func (r *Router) InitRoute(routeMap *Route) {
	r.router.HandleFunc(routeMap.Path, routeMap.Function).Name(routeMap.Name).Methods(routeMap.Method)
}

func (r *Router) Router() *mux.Router {
	return r.router
}
