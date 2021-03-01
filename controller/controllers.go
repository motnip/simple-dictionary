package controller


import (
    "fmt"
    "io/ioutil"
    "net/http"
)

type Controllers interface {
    CreateDictionary(httpResponse http.ResponseWriter, httpRequest *http.Request)
    Goodbye(httpResponse http.ResponseWriter, httpRequest *http.Request)
    Path() string
}


func  CreateDictionary(httpResponse http.ResponseWriter, httpRequest *http.Request) {
    reqBody, err := ioutil.ReadAll(httpRequest.Body)
    if err != nil {
        fmt.Fprintf(httpResponse, "Kindly enter data with the event title and description only in order to update")
    }
    input := string(reqBody[:])
    fmt.Fprintf(httpResponse, input)
}
