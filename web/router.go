package web

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Router struct {
	router *mux.Router
	paths  []*Route
}

type Route struct {
	Path           string
	ControllerFunc func(http.ResponseWriter, *http.Request)
}

func NewRouter() *Router {
	return &Router{
		router: mux.NewRouter().StrictSlash(true),
	}
}

func (r *Router) RouterStart() {
	r.router.HandleFunc(r.paths[0].Path, r.paths[0].ControllerFunc).GetMethods()
	r.router.HandleFunc(r.paths[1].Path, r.paths[1].ControllerFunc).GetMethods()
	log.Fatal(http.ListenAndServe(":8080", r.router))
}

func (r *Router) AddPath(newRoute *Route) {
	r.paths = append(r.paths, newRoute)
}
