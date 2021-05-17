package controller

import (
	"encoding/json"
	"github.com/motnip/sermo/service"
	"io/ioutil"
	"net/http"

	"github.com/motnip/sermo/model"
	"github.com/motnip/sermo/system"
	"github.com/motnip/sermo/web"
)

type WordController interface {
	AddWord(httpResponse http.ResponseWriter, httpRequest *http.Request)
	ListWords(httpResponse http.ResponseWriter, httpRequest *http.Request)
	GetAddWordRoute() *web.Route
	GetListWordRoute() *web.Route
}

type wordcontroller struct {
	wordService service.WordService
	log         *system.SermoLog
}

func NewWordController(wordService service.WordService) WordController {
	return &wordcontroller{
		wordService: wordService,
		log:         system.NewLog(),
	}
}

func (w *wordcontroller) AddWord(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	var newWordDto model.Word
	reqBody, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		w.log.LogErr(err.Error())
		http.Error(httpResponse, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &newWordDto)
	if err != nil {
		w.log.LogErr(err.Error())
		http.Error(httpResponse, "body request malformed: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = w.wordService.SaveWord(&newWordDto)
	if err != nil {
		w.log.LogErr(err.Error())
		http.Error(httpResponse, err.Error(), http.StatusBadRequest)
		return
	}

	httpResponse.WriteHeader(http.StatusCreated)

	//https://stackoverflow.com/questions/36319918/why-does-json-encoder-add-an-extra-line
	output, err := json.Marshal(newWordDto)
	if err != nil {
		w.log.LogErr(err.Error())
		http.Error(httpResponse, err.Error(), http.StatusInternalServerError)
		return
	}
	httpResponse.Write(output)
}

func (w *wordcontroller) ListWords(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	words, err := w.wordService.ListWords()
	if err != nil {
		w.log.LogErr(err.Error())
		http.Error(httpResponse, err.Error(), http.StatusBadRequest)
		return
	}
	//json.NewEncoder(httpResponse).Encode(words)
	output, err := json.Marshal(words)
	if err != nil {
		w.log.LogErr(err.Error())
		http.Error(httpResponse, err.Error(), http.StatusInternalServerError)
		return
	}
	httpResponse.Write(output)
}

func (w *wordcontroller) GetAddWordRoute() *web.Route {
	headers := make(map[string]string)
	headers["Content-type"] = "application/json"
	return &web.Route{
		Path:     "/word",
		Function: w.AddWord,
		Method:   http.MethodPost,
		Name:     "addWord",
		Headers:  &headers,
	}
}

func (w *wordcontroller) GetListWordRoute() *web.Route {
	return &web.Route{
		Path:     "/word",
		Function: w.ListWords,
		Method:   http.MethodGet,
		Name:     "listWord",
	}
}
