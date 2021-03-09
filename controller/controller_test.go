package controller

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"sermo/mocks/repository"
	"strings"
	"testing"
)

func TestCreateDictionary(t *testing.T) {
	//given
	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)

	sut := NewController(repositoryMock)

	dictionaryLanguage := "language"
	request, err := http.NewRequest("POST", "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().CreateDictionary(gomock.Eq(dictionaryLanguage)).Times(1)

	http.HandlerFunc(sut.CreateDictionary).ServeHTTP(recorder, request)

	//then
	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	responseBody := recorder.Body.String()
	if responseBody != dictionaryLanguage {
		t.Errorf("handler returned unexpected body: got %v want %v", responseBody, dictionaryLanguage)
	}

}

func TestCreateDictionary_noLanguage_returnBadRequest(t *testing.T) {
	//given
	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)

	sut := NewController(repositoryMock)

	dictionaryLanguage := ""
	request, err := http.NewRequest("POST", "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().CreateDictionary(gomock.Any()).Times(0)

	http.HandlerFunc(sut.CreateDictionary).ServeHTTP(recorder, request)

	//then
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	responseBody := recorder.Body.String()
	expectedErrorMsg := "not valid language\n"
	if responseBody != expectedErrorMsg {
		t.Errorf("handler returned unexpected body: got %v want %v", responseBody, expectedErrorMsg)
	}

}

func TestAddWord(t *testing.T) {
	//given

	newWord := "{\"Label\":\"hello\",\"Meaning\":\"ciao\",\"Sentence\":\"\"}"
	expected := "{\"Label\":\"hello\",\"Meaning\":\"ciao\",\"Sentence\":\"\"}\n"

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)
	sut := NewController(repositoryMock)

	request, err := http.NewRequest("POST", "/word", bytes.NewBuffer([]byte(newWord)))

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().AddWord(gomock.Any()).Times(1)
	http.HandlerFunc(sut.AddWord).ServeHTTP(recorder, request)

	//then
	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	responseBody := recorder.Body.String()
	if responseBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", responseBody, newWord)
	}

}

func TestAddWord_jsonMalformed_Failed(t *testing.T) {
	//given

	newWord := "{\"Label\":hello,\"Meaning\":\"ciao\",\"Sentence\":\"\"}"
	expected := "body request malformed"

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)
	sut := NewController(repositoryMock)

	request, err := http.NewRequest("POST", "/word", bytes.NewBuffer([]byte(newWord)))

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().AddWord(gomock.Any()).Times(0)
	http.HandlerFunc(sut.AddWord).ServeHTTP(recorder, request)

	//then
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	responseBody := recorder.Body.String()
	if !strings.Contains(responseBody,expected){
		t.Errorf("handler returned unexpected body: got %v want %v", responseBody, expected)
	}

}
