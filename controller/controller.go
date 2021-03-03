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

func NewController(repository model.Repository) Controllers{
	return &controller{
		repository: repository,
	}
}
func (c *controller)CreateDictionary(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	reqBody, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		fmt.Fprintf(httpResponse, "Kindly enter data with the event title and description only in order to update")
	}
	httpResponse.WriteHeader(http.StatusCreated)
	input := string(reqBody[:])

	c.repository.CreateDictionary(input)

	fmt.Fprintf(httpResponse, input)
}
