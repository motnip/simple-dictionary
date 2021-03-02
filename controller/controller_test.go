package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDictionary(t *testing.T) {

	dictionaryLanguage := "language"
	request, err := http.NewRequest("GET", "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	http.HandlerFunc(CreateDictionary).ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	responseBody := recorder.Body.String()
	if responseBody != dictionaryLanguage {
		t.Errorf("handler returned unexpected body: got %v want %v", responseBody, dictionaryLanguage)
	}

}
