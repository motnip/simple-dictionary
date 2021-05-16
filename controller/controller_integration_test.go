package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/motnip/sermo/model"
	"github.com/motnip/sermo/web"
)

var router *mux.Router
var repository model.Repository
var dictionaryController DictionaryController
var wordController WordController

func TestMain(m *testing.M) {
	setUp()
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestIntegration_Controller_AddNewWord_Succeed(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	defer tearDown()

	//given
	dictionaryLanguage := "en"
	newWord := "{\"Label\":\"hello\",\"Meaning\":\"ciao\",\"Sentence\":\"\"}"

	expectedWordsList := "[" + newWord + "]"

	requestCreateDictionary, err := http.NewRequest(http.MethodPost, "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	requestCreateDictionary.Header.Add("Content-type", "application/json")

	if err != nil {
		t.Fatal(err)
	}
	requestAddNewWord, err := http.NewRequest(http.MethodPost, "/word", bytes.NewBuffer([]byte(newWord)))
	requestAddNewWord.Header.Add("Content-type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	requestListAllWord, err := http.NewRequest(http.MethodGet, "/word", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorderCreateDictionary := httptest.NewRecorder()
	recorderAddWord := httptest.NewRecorder()
	recorderListWord := httptest.NewRecorder()

	//when
	router.ServeHTTP(recorderCreateDictionary, requestCreateDictionary)
	router.ServeHTTP(recorderAddWord, requestAddNewWord)
	router.ServeHTTP(recorderListWord, requestListAllWord)

	//then
	if status := recorderCreateDictionary.Code; status != http.StatusCreated {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	if status := recorderAddWord.Code; status != http.StatusCreated {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
	responseBody := recorderAddWord.Body.String()
	if responseBody != newWord {
		t.Errorf("Router returned unexpected body: got %v want %v", responseBody, newWord)
	}
	if status := recorderListWord.Code; status != http.StatusOK {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	responseBody = recorderListWord.Body.String()
	if responseBody != expectedWordsList {
		t.Errorf("Router returned unexpected body: got %v want %v", responseBody, expectedWordsList)
	}
}

func TestIntegration_Controller_CreateDictionary_Succeed(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	defer tearDown()

	//given
	dictionaryLanguage := "en"

	requestCreateDictionary, err := http.NewRequest(http.MethodPost, "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	requestCreateDictionary.Header.Add("Content-type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	requestListDictionary, err := http.NewRequest(http.MethodGet, "/dictionary", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorderCreateDictionary := httptest.NewRecorder()
	recorderListDictionary := httptest.NewRecorder()

	expectedList := "[{\"Language\":\"en\",\"Words\":null}]"
	//when
	router.ServeHTTP(recorderCreateDictionary, requestCreateDictionary)
	router.ServeHTTP(recorderListDictionary, requestListDictionary)

	//then
	if status := recorderCreateDictionary.Code; status != http.StatusCreated {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	if status := recorderListDictionary.Code; status != http.StatusOK {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	responseBody := recorderListDictionary.Body.String()
	if responseBody != expectedList {
		t.Errorf("Router returned unexpected body: got %v want %v", responseBody, expectedList)
	}
}

func TestIntegration_Controller_NoDictionaryExists_Failed(t *testing.T) {

	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	defer tearDown()

	//given
	newWord := "{\"Label\":\"hello\",\"Meaning\":\"ciao\",\"Sentence\":\"\"}"
	expectedErrorMessage := "no dictionary available\n"

	requestAddNewWord, err := http.NewRequest(http.MethodPost, "/word", bytes.NewBuffer([]byte(newWord)))
	requestAddNewWord.Header.Add("Content-type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	requestListAllWord, err := http.NewRequest(http.MethodGet, "/word", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorderAddWord := httptest.NewRecorder()
	recorderListWord := httptest.NewRecorder()

	//when
	router.ServeHTTP(recorderAddWord, requestAddNewWord)
	router.ServeHTTP(recorderListWord, requestListAllWord)

	//then

	if status := recorderAddWord.Code; status != http.StatusBadRequest {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
	responseBody := recorderAddWord.Body.String()
	if responseBody != expectedErrorMessage {
		t.Errorf("Router returned unexpected body: got %v want %v", responseBody, expectedErrorMessage)
	}

	if status := recorderListWord.Code; status != http.StatusBadRequest {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	responseBody = recorderListWord.Body.String()
	if responseBody != expectedErrorMessage {
		t.Errorf("Router returned unexpected body: got %v want %v", responseBody, expectedErrorMessage)
	}
}

func setUp() {
	repository = model.NewRepository()
	dictionaryController = NewController(repository)
	wordController = NewWordController(repository)

	newRouter := web.NewRouter()
	newRouter.InitRoute(dictionaryController.GetCreateDictionaryRoute())
	newRouter.InitRoute(dictionaryController.GetListAllDictionary())
	newRouter.InitRoute(wordController.GetAddWordRoute())
	newRouter.InitRoute(wordController.GetListWordRoute())

	router = newRouter.Router()
}

func tearDown() {
	repository.DeleteDictionary()
}
