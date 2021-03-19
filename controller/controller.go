package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sermo/model"
)

type Controller interface {
	CreateDictionary(httpResponse http.ResponseWriter, httpRequest *http.Request)
	AddWord(httpResponse http.ResponseWriter, httpRequest *http.Request)
	ListWords(httpResponse http.ResponseWriter, httpRequest *http.Request)
}

type controller struct {
	repository model.Repository
}

func NewController(repository model.Repository) Controller {
	return &controller{
		repository: repository,
	}
}

func (c *controller) CreateDictionary(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	var input string
	reqBody, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		http.Error(httpResponse, err.Error(), http.StatusBadRequest)
	}

	if input = string(reqBody[:]); input == "" {
		http.Error(httpResponse, "not valid language", http.StatusBadRequest)
		return
	}

	_, err = c.repository.CreateDictionary(input)

	if err != nil {
		http.Error(httpResponse, err.Error(), http.StatusBadRequest)
		return
	}

	httpResponse.WriteHeader(http.StatusCreated)
	httpResponse.Write([]byte(input))
}

func (c *controller) AddWord(httpResponse http.ResponseWriter, httpRequest *http.Request) {
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

	err = c.repository.AddWord(&newWordDto)
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

func (c *controller) ListWords(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	words, err := c.repository.ListWords()
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
