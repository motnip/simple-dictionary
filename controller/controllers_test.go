package controller

import (
    "bytes"
    "net/http/httptest"
    "testing"
    "net/http"
)

func TestCreateDictionary(t *testing.T) {

    dictionaryLanguage := "language"
    //request, err := http.Post("/dictionary", "text/string", bytes.NewBuffer([]byte(dictionaryLanguage)))
    request, err := http.NewRequest("GET", "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
    if err != nil {
        t.Fatal(err)
    }

    recorder := httptest.NewRecorder()
    http.HandlerFunc(CreateDictionary).ServeHTTP(recorder, request)

    if status := recorder.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",status, http.StatusOK)
    }

    // Check the response body is what we expect.

    responseBody := recorder.Body.String()
    if responseBody != dictionaryLanguage {
        t.Errorf("handler returned unexpected body: got %v want %v",responseBody, dictionaryLanguage)
    }

}
