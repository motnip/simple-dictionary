package web

import "net/http"

type Route struct {
	Path     string
	Function func(http.ResponseWriter, *http.Request)
	Method   string
	Name     string
}

/*
type RoutMap struct {
	Handler []*Route
}

func (h *RoutMap) Add(route *Route) {
	//TODO add check on Route fields
	h.Handler = append(h.Handler, route)
}

func (h *RoutMap) AllRoutes() []*Route {
	return h.Handler
}
*/
