package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sermo/controller"
)

type Router struct {
	router *mux.Router
	paths  []*route
}

type route struct {
	path           string
	controllerFunc func(http.ResponseWriter, *http.Request)
}

func NewRouter() *Router {
	return &Router{
		router: mux.NewRouter().StrictSlash(true),
	}
}

func (r *Router) RouterStart() {
	if len(r.paths) == 0 {
		//TODO raise a panic!!!
		//return errors.New("no path available")
		fmt.Println("Welcome home!")
	}
	//  r.router.HandleFunc(r.paths[0].path, r.paths[0].controllerFunc).GetMethods()
	//   log.Fatal(http.ListenAndServe(":8080", r.router))
}

func (r *Router) AddPath(c controller.Controller) {
	r.paths = append(r.paths, &route{
		path:           c.Path(),
		controllerFunc: c.Listen,
	})
}
