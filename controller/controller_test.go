package controller

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"sermo/mocks/repository"
	"sermo/model"
	"strings"
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
	router.HandleFunc("/dictionary", sut.CreateDictionary).Methods(http.MethodPost)

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
	router.HandleFunc("/dictionary", sut.CreateDictionary).Methods(http.MethodPost)

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
	router.HandleFunc("/dictionary", sut.CreateDictionary).Methods(http.MethodPost)

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

func TestAddWord(t *testing.T) {
	//given
	newWord := "{\"Label\":\"hello\",\"Meaning\":\"ciao\",\"Sentence\":\"\"}"
	returnedWord := "{\"Label\":\"hello\",\"Meaning\":\"ciao\",\"Sentence\":\"\"}"

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)
	sut := NewController(repositoryMock)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/word", sut.AddWord).Methods(http.MethodPost)

	request, err := http.NewRequest(http.MethodPost, "/word", bytes.NewBuffer([]byte(newWord)))

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().AddWord(gomock.Any()).Times(1)
	router.ServeHTTP(recorder, request)

	//then
	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	responseBody := recorder.Body.String()
	if responseBody != returnedWord {
		t.Errorf("Router returned unexpected body: got %v want %v", responseBody, returnedWord)
	}

}

func TestAddWord_noDictionary_Failed(t *testing.T) {
	//given
	newWord := "{\"Label\":\"hello\",\"Meaning\":\"ciao\",\"Sentence\":\"\"}"
	expectedError := errors.New("no dictionary available")

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)
	sut := NewController(repositoryMock)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/word", sut.AddWord).Methods(http.MethodPost)

	request, err := http.NewRequest(http.MethodPost, "/word", bytes.NewBuffer([]byte(newWord)))

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().AddWord(gomock.Any()).Return(expectedError)
	router.ServeHTTP(recorder, request)

	//then
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	responseBody := recorder.Body.String()
	if !strings.Contains(responseBody, expectedError.Error()) {
		t.Errorf("Router returned unexpected body: got %v want %v", responseBody, expectedError)
	}

}

func TestAddWord_jsonMalformed_Failed(t *testing.T) {
	//given
	newWord := "{\"Label\":hello,\"Meaning\":\"ciao\",\"Sentence\":\"\"}"
	expectedErrorMessage := "body request malformed"

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)
	sut := NewController(repositoryMock)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/word", sut.AddWord).Methods(http.MethodPost)

	request, err := http.NewRequest(http.MethodPost, "/word", bytes.NewBuffer([]byte(newWord)))

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().AddWord(gomock.Any()).Times(0)
	router.ServeHTTP(recorder, request)

	//then
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	responseBody := recorder.Body.String()
	if !strings.Contains(responseBody, expectedErrorMessage) {
		t.Errorf("Router returned unexpected body: got %v want %v", responseBody, expectedErrorMessage)
	}

}

func TestListWords(t *testing.T) {
	//given
	words := make([]*model.Word, 0)
	words = append(words, &model.Word{
		Label:   "foo",
		Meaning: "foo",
	})
	words = append(words, &model.Word{
		Label:   "bar",
		Meaning: "bar",
	})

	expectedWordsList := "[{\"Label\":\"foo\",\"Meaning\":\"foo\",\"Sentence\":\"\"},{\"Label\":\"bar\",\"Meaning\":\"bar\",\"Sentence\":\"\"}]"

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)
	sut := NewController(repositoryMock)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/word", sut.ListWords).Methods(http.MethodGet)

	request, err := http.NewRequest(http.MethodGet, "/word", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().ListWords().Return(words, nil)
	router.ServeHTTP(recorder, request)

	//then
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	responseBody := recorder.Body.String()
	if responseBody != expectedWordsList {
		t.Errorf("Router returned unexpected body: got %v want %v", responseBody, expectedWordsList)
	}

}

func TestListWords_noAvailableDictionary_returnBadRequest(t *testing.T) {
	//given
	expectedError := errors.New("no dictionary available")

	controller := gomock.NewController(t)
	repositoryMock := mock_model.NewMockRepository(controller)
	sut := NewController(repositoryMock)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/word", sut.ListWords).Methods(http.MethodGet)

	request, err := http.NewRequest(http.MethodGet, "/word", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	//when
	repositoryMock.EXPECT().ListWords().Return(nil, expectedError)
	router.ServeHTTP(recorder, request)

	//then
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	responseBody := recorder.Body.String()
	if !strings.Contains(responseBody, expectedError.Error()) {
		t.Errorf("Router returned unexpected body: got %v want %v", responseBody, expectedError.Error())
	}
}
