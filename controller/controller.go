package controller

import (
	"net/http"
)

type Controller interface {
	Greetings(httpResponse http.ResponseWriter, httpRequest *http.Request)
	Goodbye(httpResponse http.ResponseWriter, httpRequest *http.Request)
	Path() string
}

type route struct {
	path string
}
