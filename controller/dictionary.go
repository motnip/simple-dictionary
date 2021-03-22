package controller

import (
	"io/ioutil"
	"net/http"
	"sermo/model"
	"sermo/web"
)

type DictionaryController interface {
	CreateDictionary(httpResponse http.ResponseWriter, httpRequest *http.Request)
	GetCreateDictionaryRoute() *web.Route
}

type dictionarycontroller struct {
	repository model.Repository
}

func NewController(repository model.Repository) DictionaryController {
	return &dictionarycontroller{
		repository: repository,
	}
}

func (d *dictionarycontroller) CreateDictionary(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	var input string
	reqBody, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		http.Error(httpResponse, err.Error(), http.StatusBadRequest)
	}

	if input = string(reqBody[:]); input == "" {
		http.Error(httpResponse, "not valid language", http.StatusBadRequest)
		return
	}

	_, err = d.repository.CreateDictionary(input)

	if err != nil {
		http.Error(httpResponse, err.Error(), http.StatusBadRequest)
		return
	}

	httpResponse.WriteHeader(http.StatusCreated)
	httpResponse.Write([]byte(input))
}

func (d *dictionarycontroller) GetCreateDictionaryRoute() *web.Route {
	return &web.Route{
		Path:     "/dictionary",
		Function: d.CreateDictionary,
		Method:   http.MethodPost,
		Name:     "createDictionary",
	}
}
