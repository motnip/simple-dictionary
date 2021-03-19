package controller

import (
	"bytes"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"os"
	"sermo/model"
	"sermo/web"
	"testing"
)

var router *mux.Router

func TestMain(m *testing.M) {

	repository := model.NewRepository()
	controller := NewController(repository)

	newRouter := web.NewRouter()
	newRouter.InitRoute(&web.Route{
		Path:     "/dictionary",
		Function: controller.CreateDictionary,
		Method:   http.MethodPost,
		Name:     "createDictionary",
	})
	newRouter.InitRoute(&web.Route{
		Path:     "/word",
		Function: controller.AddWord,
		Method:   http.MethodPost,
		Name:     "addWord",
	})
	newRouter.InitRoute(&web.Route{
		Path:     "/word",
		Function: controller.ListWords,
		Method:   http.MethodGet,
		Name:     "listWords",
	})

	router = newRouter.Router()

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestController_AddNewWord_Succeed(t *testing.T) {

	//given
	dictionaryLanguage := "en"
	newWord := "{\"Label\":\"hello\",\"Meaning\":\"ciao\",\"Sentence\":\"\"}"

	expectedWordsList := "[" + newWord + "]"

	requestCreateDictionary, err := http.NewRequest(http.MethodPost, "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	if err != nil {
		t.Fatal(err)
	}
	requestAddNewWord, err := http.NewRequest(http.MethodPost, "/word", bytes.NewBuffer([]byte(newWord)))
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
