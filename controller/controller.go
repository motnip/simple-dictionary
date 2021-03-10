package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sermo/model"
)

//to rename Controller
type Controllers interface {
	CreateDictionary(httpResponse http.ResponseWriter, httpRequest *http.Request)
	AddWord(httpResponse http.ResponseWriter, httpRequest *http.Request)
}

type controller struct {
	repository model.Repository
}

func NewController(repository model.Repository) Controllers {
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
	fmt.Fprintf(httpResponse, input)
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

	c.repository.AddWord(&newWordDto)

	httpResponse.WriteHeader(http.StatusCreated)

	//https://stackoverflow.com/questions/36319918/why-does-json-encoder-add-an-extra-line
	json.NewEncoder(httpResponse).Encode(newWordDto)
}
