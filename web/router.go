package web

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	c "sermo/controller"
	"sermo/model"
)

type Router struct {
	router *mux.Router
	controller c.Controllers
}

func NewRouter() *Router {
	return &Router{
		router: mux.NewRouter().StrictSlash(true),
		controller: c.NewController(model.NewRepository()),
	}
}

func (r *Router) RouterStart() {
	r.router.HandleFunc("/dictionary", r.controller.CreateDictionary)
	r.router.HandleFunc("/word", r.controller.AddWord)
	r.router.HandleFunc("/word", r.controller.ListWords).GetMethods()
	log.Fatal(http.ListenAndServe(":8080", r.router))
}
