package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sermo/model"
)

//to rename Controller
type Controllers interface {
	CreateDictionary(httpResponse http.ResponseWriter, httpRequest *http.Request)
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
		fmt.Fprintf(httpResponse, "Kindly enter data with the event title and description only in order to update")
	}

	if input = string(reqBody[:]); input == "" {
		httpResponse.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(httpResponse, input)
		return
	}

	httpResponse.WriteHeader(http.StatusCreated)
	c.repository.CreateDictionary(input)
	fmt.Fprintf(httpResponse, input)
}
