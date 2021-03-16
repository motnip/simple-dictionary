package controller

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"sermo/mocks/repository"
	"sermo/model"
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

func TestCreateDictionary_existDictionaryForALanguage_returnBadRequest(t *testing.T) {
	//given
	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)

	sut := NewController(repositoryMock)

	dictionaryLanguage := "en"
	firstRequest, err := http.NewRequest("POST", "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	if err != nil {
		t.Fatal(err)
	}
	secondRequest, err := http.NewRequest("POST", "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	if err != nil {
		t.Fatal(err)
	}

	recorderFirstRequest := httptest.NewRecorder()
	recorderSecondRequest := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().CreateDictionary(gomock.Any()).Return(nil, nil)
	repositoryMock.EXPECT().CreateDictionary(gomock.Any()).Return(nil, errors.New("forced error"))
	http.HandlerFunc(sut.CreateDictionary).ServeHTTP(recorderFirstRequest, firstRequest)
	http.HandlerFunc(sut.CreateDictionary).ServeHTTP(recorderSecondRequest, secondRequest)

	//then
	if status := recorderFirstRequest.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	if status := recorderSecondRequest.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
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

func TestAddWord_noDictionary_Failed(t *testing.T) {
	//given

	newWord := "{\"Label\":\"hello\",\"Meaning\":\"ciao\",\"Sentence\":\"\"}"
	expectedError := errors.New("no dictionary available")

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)
	sut := NewController(repositoryMock)

	request, err := http.NewRequest("POST", "/word", bytes.NewBuffer([]byte(newWord)))

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().AddWord(gomock.Any()).Return(expectedError)
	http.HandlerFunc(sut.AddWord).ServeHTTP(recorder, request)

	//then
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	responseBody := recorder.Body.String()
	if !strings.Contains(responseBody, expectedError.Error()) {
		t.Errorf("handler returned unexpected body: got %v want %v", responseBody, expectedError)
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
	if !strings.Contains(responseBody, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", responseBody, expected)
	}

}

func TestListWords(t *testing.T) {
	//given
	controller := gomock.NewController(t)

	words := make([]*model.Word, 0)
	words = append(words, &model.Word{
		Label:   "foo",
		Meaning: "foo",
	})
	words = append(words, &model.Word{
		Label:   "bar",
		Meaning: "bar",
	})

	expected := "[{\"Label\":\"foo\",\"Meaning\":\"foo\",\"Sentence\":\"\"},{\"Label\":\"bar\",\"Meaning\":\"bar\",\"Sentence\":\"\"}]\n"

	repositoryMock := mock_model.NewMockRepository(controller)
	sut := NewController(repositoryMock)

	request, err := http.NewRequest("get", "/word", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().ListWords().Return(words, nil)
	http.HandlerFunc(sut.ListWords).ServeHTTP(recorder, request)

	//then
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	responseBody := recorder.Body.String()
	if responseBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", responseBody, expected)
	}

}

func TestListWords_noDictionaryForALanguage_returnBadRequest(t *testing.T) {
	//given
	controller := gomock.NewController(t)

	expected := "no dictionary available"
	expectedError := errors.New(expected)

	repositoryMock := mock_model.NewMockRepository(controller)
	sut := NewController(repositoryMock)

	request, err := http.NewRequest("get", "/word", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().ListWords().Return(nil, expectedError)
	http.HandlerFunc(sut.ListWords).ServeHTTP(recorder, request)

	//then
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	responseBody := recorder.Body.String()
	if !strings.Contains(responseBody, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", responseBody, expected)
	}
}
