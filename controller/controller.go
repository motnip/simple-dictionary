package controller

import (
	"net/http"
)

type Controller interface {
	Listen(httpResponse http.ResponseWriter, httpRequest *http.Request)
	Path() string
}

type route struct {
	path string
}
