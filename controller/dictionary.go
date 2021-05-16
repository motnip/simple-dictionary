package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/motnip/sermo/model"
	"github.com/motnip/sermo/system"
	"github.com/motnip/sermo/web"
)

type DictionaryController interface {
	CreateDictionary(httpResponse http.ResponseWriter, httpRequest *http.Request)
	ListAllDictionary(httpResponse http.ResponseWriter, httpRequest *http.Request)
	GetCreateDictionaryRoute() *web.Route
	GetListAllDictionary() *web.Route
}

type dictionarycontroller struct {
	repository model.Repository
	log        *system.SermoLog
}

func NewController(repository model.Repository) DictionaryController {
	return &dictionarycontroller{
		repository: repository,
		log:        system.NewLog(),
	}
}

func (d *dictionarycontroller) CreateDictionary(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	var input string
	reqBody, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		d.log.LogErr(err.Error())
		http.Error(httpResponse, err.Error(), http.StatusBadRequest)
	}

	if input = string(reqBody[:]); input == "" {
		d.log.LogErr("Language cannot be empty string")
		http.Error(httpResponse, "not valid language", http.StatusBadRequest)
		return
	}

	_, err = d.repository.CreateDictionary(input)

	if err != nil {
		d.log.LogErr(err.Error())
		http.Error(httpResponse, err.Error(), http.StatusBadRequest)
		return
	}
	d.log.LogInfo(fmt.Sprintf("Dicitionary for language %s created successfuly", input))
	httpResponse.WriteHeader(http.StatusCreated)
	httpResponse.Write([]byte(input))
}

func (d *dictionarycontroller) ListAllDictionary(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	dictionaries := d.repository.ListDictionary()

	output, err := json.Marshal(dictionaries)
	if err != nil {
		d.log.LogErr(err.Error())
		http.Error(httpResponse, err.Error(), http.StatusInternalServerError)
		return
	}
	httpResponse.WriteHeader(http.StatusOK)
	httpResponse.Write(output)
}

func (d *dictionarycontroller) GetCreateDictionaryRoute() *web.Route {
	headers := make(map[string]string)
	headers["Content-type"] = "application/json"
	return &web.Route{
		Path:     "/dictionary",
		Function: d.CreateDictionary,
		Method:   http.MethodPost,
		Name:     "createDictionary",
		Headers:  &headers,
	}
}

func (d *dictionarycontroller) GetListAllDictionary() *web.Route {
	return &web.Route{
		Path:     "/dictionary",
		Function: d.ListAllDictionary,
		Method:   http.MethodGet,
		Name:     "listAllDictionary",
	}
}
