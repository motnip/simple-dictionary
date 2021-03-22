package controller

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"sermo/mocks/repository"
	"testing"
)

func TestCreateDictionary(t *testing.T) {
	//given
	dictionaryLanguage := "language"

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)
	recorder := httptest.NewRecorder()
	sut := NewController(repositoryMock)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(sut.GetCreateDictionaryRoute().Path, sut.GetCreateDictionaryRoute().Function).Methods(sut.GetCreateDictionaryRoute().Method)

	request, err := http.NewRequest(http.MethodPost, "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	if err != nil {
		t.Fatal(err)
	}

	//when
	repositoryMock.EXPECT().CreateDictionary(gomock.Eq(dictionaryLanguage)).Times(1)
	router.ServeHTTP(recorder, request)

	//then
	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	responseBody := recorder.Body.String()
	if responseBody != dictionaryLanguage {
		t.Errorf("Router returned unexpected body: got %v want %v", responseBody, dictionaryLanguage)
	}

}

func TestCreateDictionary_EmptyLanguage_returnBadRequest(t *testing.T) {
	//given
	emptyDictionaryLanguage := ""

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)
	sut := NewController(repositoryMock)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(sut.GetCreateDictionaryRoute().Path, sut.GetCreateDictionaryRoute().Function).Methods(sut.GetCreateDictionaryRoute().Method)

	request, err := http.NewRequest(http.MethodPost, "/dictionary", bytes.NewBuffer([]byte(emptyDictionaryLanguage)))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().CreateDictionary(gomock.Any()).Times(0)
	router.ServeHTTP(recorder, request)

	//then
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	responseBody := recorder.Body.String()
	expectedErrorMsg := "not valid language\n"
	if responseBody != expectedErrorMsg {
		t.Errorf("Router returned unexpected body: got %v want %v", responseBody, expectedErrorMsg)
	}
}

func TestCreateDictionary_existDictionaryForALanguage_returnBadRequest(t *testing.T) {
	//given
	dictionaryLanguage := "en"

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)
	sut := NewController(repositoryMock)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(sut.GetCreateDictionaryRoute().Path, sut.GetCreateDictionaryRoute().Function).Methods(sut.GetCreateDictionaryRoute().Method)

	firstRequest, err := http.NewRequest(http.MethodPost, "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	if err != nil {
		t.Fatal(err)
	}
	secondRequest, err := http.NewRequest(http.MethodPost, "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	if err != nil {
		t.Fatal(err)
	}

	recorderFirstRequest := httptest.NewRecorder()
	recorderSecondRequest := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().CreateDictionary(gomock.Any()).Return(nil, nil)
	repositoryMock.EXPECT().CreateDictionary(gomock.Any()).Return(nil, errors.New("forced error"))
	router.ServeHTTP(recorderFirstRequest, firstRequest)
	router.ServeHTTP(recorderSecondRequest, secondRequest)

	//then
	if status := recorderFirstRequest.Code; status != http.StatusCreated {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	if status := recorderSecondRequest.Code; status != http.StatusBadRequest {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
