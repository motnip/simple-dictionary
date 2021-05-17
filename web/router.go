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

type Route struct {
	Path     string
	Function func(http.ResponseWriter, *http.Request)
	Method   string
	Name     string
	Headers  *map[string]string
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
	route := r.router.HandleFunc(routeMap.Path, routeMap.Function)
	route.Name(routeMap.Name)
	route.Methods(routeMap.Method)
	if routeMap.Headers != nil && len(*routeMap.Headers) > 0 {
		route.Headers("Content-type", "application/json")
	}
}

func (r *Router) Router() *mux.Router {
	return r.router
}
