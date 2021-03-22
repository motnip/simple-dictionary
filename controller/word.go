package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sermo/model"
)

type WordController interface {
	AddWord(httpResponse http.ResponseWriter, httpRequest *http.Request)
	ListWords(httpResponse http.ResponseWriter, httpRequest *http.Request)
}

type wordcontroller struct {
	repository model.Repository
}

func NewWordController(repository model.Repository) WordController {
	return &wordcontroller{
		repository: repository,
	}
}

func (w *wordcontroller) AddWord(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	var newWordDto model.Word
	reqBody, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		http.Error(httpResponse, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &newWordDto)
	if err != nil {
		http.Error(httpResponse, "body request malformed: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = w.repository.AddWord(&newWordDto)
	if err != nil {
		http.Error(httpResponse, err.Error(), http.StatusBadRequest)
		return
	}

	httpResponse.WriteHeader(http.StatusCreated)

	//https://stackoverflow.com/questions/36319918/why-does-json-encoder-add-an-extra-line
	output, err := json.Marshal(newWordDto)
	if err != nil {
		http.Error(httpResponse, err.Error(), http.StatusInternalServerError)
		return
	}
	httpResponse.Write(output)
}

func (w *wordcontroller) ListWords(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	words, err := w.repository.ListWords()
	if err != nil {
		http.Error(httpResponse, err.Error(), http.StatusBadRequest)
		return
	}
	//json.NewEncoder(httpResponse).Encode(words)
	output, err := json.Marshal(words)
	if err != nil {
		http.Error(httpResponse, err.Error(), http.StatusInternalServerError)
		return
	}
	httpResponse.Write(output)
}
