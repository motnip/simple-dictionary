package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

func NewRouter() *Router {
	return &Router{
		router: mux.NewRouter().StrictSlash(true),
	}
}

func (r *Router) RouterStart() {
	fmt.Println("Server started... ")
	log.Fatal(http.ListenAndServe(":3000", r.router))
}

func (r *Router) InitRoute(routeMap *Route) {
	r.router.HandleFunc(routeMap.Path, routeMap.Function).Name(routeMap.Name).Methods(routeMap.Method)
}

func (r *Router) Router() *mux.Router {
	return r.router
}
