package controller

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"sermo/mocks/repository"
)

func TestCreateDictionary(t *testing.T) {
	//given
	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)

	sut := NewController(repositoryMock)

	dictionaryLanguage := "language"
	request, err := http.NewRequest("GET", "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().CreateDictionary(gomock.Any()).Times(1)

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
