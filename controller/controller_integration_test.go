package controller

import (
	"bytes"
	"github.com/motnip/sermo/service"
	"github.com/motnip/sermo/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/motnip/sermo/model"
)

var router *mux.Router
var repository model.Repository
var wordService service.WordService
var dictionaryController DictionaryController
var wordController WordController

type IntegrationTestSuite struct {
	suite.Suite
	repository           model.Repository
	wordService          service.WordService
	dictionaryController DictionaryController
	wordController       WordController
	newRouter            *web.Router
	router               *mux.Router
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (suite *IntegrationTestSuite) SetupSuite() {
	suite.repository = model.NewRepository()
	suite.wordService = service.NewWordService(suite.repository)
	suite.dictionaryController = NewController(suite.repository)
	suite.wordController = NewWordController(suite.wordService)

	suite.newRouter = web.NewRouter()
	suite.newRouter.InitRoute(suite.dictionaryController.GetCreateDictionaryRoute())
	suite.newRouter.InitRoute(suite.dictionaryController.GetListAllDictionary())
	suite.newRouter.InitRoute(suite.wordController.GetAddWordRoute())
	suite.newRouter.InitRoute(suite.wordController.GetListWordRoute())

	suite.router = suite.newRouter.Router()
}

func (suite *IntegrationTestSuite) TestIntegration_Controller_AddNewWord_Succeed() {

	if testing.Short() {
		suite.T().Skip("skipping testing in short mode")
	}

	//given
	dictionaryLanguage := "en"
	newWord := "{\"Label\":\"hello\",\"Meaning\":\"ciao\",\"Sentence\":\"\"}"
	expectedWordsList := "[" + newWord + "]"

	requestCreateDictionary, err := http.NewRequest(http.MethodPost, "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	requestCreateDictionary.Header.Add("Content-type", "application/json")
	suite.Require().NoError(err)

	requestAddNewWord, err := http.NewRequest(http.MethodPost, "/word", bytes.NewBuffer([]byte(newWord)))
	requestAddNewWord.Header.Add("Content-type", "application/json")
	suite.Require().NoError(err)

	requestListAllWord, err := http.NewRequest(http.MethodGet, "/word", nil)
	suite.Require().NoError(err)

	recorderCreateDictionary := httptest.NewRecorder()
	recorderAddWord := httptest.NewRecorder()
	recorderListWord := httptest.NewRecorder()

	//when
	suite.router.ServeHTTP(recorderCreateDictionary, requestCreateDictionary)
	suite.router.ServeHTTP(recorderAddWord, requestAddNewWord)
	suite.router.ServeHTTP(recorderListWord, requestListAllWord)

	//then
	assert.Equal(suite.T(), recorderCreateDictionary.Code, http.StatusCreated)
	assert.Equal(suite.T(), recorderAddWord.Code, http.StatusCreated)
	assert.Equal(suite.T(), recorderAddWord.Body.String(), newWord)
	assert.Equal(suite.T(), recorderListWord.Code, http.StatusOK)
	assert.Equal(suite.T(), recorderListWord.Body.String(), expectedWordsList)
}

func (suite *IntegrationTestSuite) TestIntegration_Controller_CreateDictionary_Succeed() {

	if testing.Short() {
		suite.T().Skip("skipping testing in short mode")
	}

	//given
	dictionaryLanguage := "en"
	requestCreateDictionary, err := http.NewRequest(http.MethodPost, "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	requestCreateDictionary.Header.Add("Content-type", "application/json")
	suite.Require().NoError(err)

	requestListDictionary, err := http.NewRequest(http.MethodGet, "/dictionary", nil)
	suite.Require().NoError(err)

	recorderCreateDictionary := httptest.NewRecorder()
	recorderListDictionary := httptest.NewRecorder()
	expectedList := "[{\"Language\":\"en\",\"Words\":null}]"

	//when
	suite.router.ServeHTTP(recorderCreateDictionary, requestCreateDictionary)
	suite.router.ServeHTTP(recorderListDictionary, requestListDictionary)

	//then
	assert.Equal(suite.T(), recorderCreateDictionary.Code, http.StatusCreated)
	assert.Equal(suite.T(), recorderListDictionary.Code, http.StatusOK)
	assert.Equal(suite.T(), recorderListDictionary.Body.String(), expectedList)
}

func (suite *IntegrationTestSuite) TestIntegration_Controller_NoDictionaryExists_Failed() {

	if testing.Short() {
		suite.T().Skip("skipping testing in short mode")
	}

	//given
	newWord := "{\"Label\":\"hello\",\"Meaning\":\"ciao\",\"Sentence\":\"\"}"
	expectedErrorMessage := "no dictionary available\n"
	requestAddNewWord, err := http.NewRequest(http.MethodPost, "/word", bytes.NewBuffer([]byte(newWord)))
	requestAddNewWord.Header.Add("Content-type", "application/json")
	suite.Require().NoError(err)

	requestListAllWord, err := http.NewRequest(http.MethodGet, "/word", nil)
	suite.Require().NoError(err)

	recorderAddWord := httptest.NewRecorder()
	recorderListWord := httptest.NewRecorder()

	//when
	suite.router.ServeHTTP(recorderAddWord, requestAddNewWord)
	suite.router.ServeHTTP(recorderListWord, requestListAllWord)

	//then
	assert.Equal(suite.T(), recorderAddWord.Code, http.StatusBadRequest)
	assert.Equal(suite.T(), recorderAddWord.Body.String(), expectedErrorMessage)
	assert.Equal(suite.T(), recorderListWord.Code, http.StatusBadRequest)
	assert.Equal(suite.T(), recorderListWord.Body.String(), expectedErrorMessage)
}

func (suite *IntegrationTestSuite) TearDownTest() {
	suite.repository.DeleteDictionary()
}
