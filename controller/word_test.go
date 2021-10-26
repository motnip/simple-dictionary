package controller

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	mock_service "github.com/motnip/sermo/mocks/service"
	"github.com/motnip/sermo/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type WordTestSuite struct {
	suite.Suite
	controller  *gomock.Controller
	serviceMock *mock_service.MockWordService
	router      *mux.Router
	sut         WordController
}

func TestWordTestSuite(t *testing.T) {
	suite.Run(t, new(WordTestSuite))
}

func (w *WordTestSuite) SetupTest() {
	w.controller = gomock.NewController(w.T())
	w.serviceMock = mock_service.NewMockWordService(w.controller)
	w.sut = NewWordController(w.serviceMock)
	w.router = mux.NewRouter().StrictSlash(true)
	w.router.HandleFunc(w.sut.GetAddWordRoute().Path, w.sut.GetAddWordRoute().Function).Methods(w.sut.GetAddWordRoute().Method)
	w.router.HandleFunc(w.sut.GetListWordRoute().Path, w.sut.GetListWordRoute().Function).Methods(w.sut.GetListWordRoute().Method)
}

func (w *WordTestSuite) TestAddWord() {
	//given
	newWord := "{\"Label\":\"hello\",\"Meaning\":\"ciao\",\"Sentence\":\"\"}"
	returnedWord := "{\"Label\":\"hello\",\"Meaning\":\"ciao\",\"Sentence\":\"\"}"

	request, err := http.NewRequest(http.MethodPost, "/word", bytes.NewBuffer([]byte(newWord)))
	w.Require().NoError(err)

	recorder := httptest.NewRecorder()

	//when
	w.serviceMock.EXPECT().SaveWord(gomock.Any()).Times(1)
	w.router.ServeHTTP(recorder, request)

	//then
	assert.Equal(w.T(), recorder.Code, http.StatusCreated)
	assert.Equal(w.T(), returnedWord, recorder.Body.String())
}

func (w *WordTestSuite) TestAddWord_noDictionary_Failed() {
	//given
	newWord := "{\"Label\":\"hello\",\"Meaning\":\"ciao\",\"Sentence\":\"\"}"
	expectedError := errors.New("no dictionary available")

	request, err := http.NewRequest(http.MethodPost, "/word", bytes.NewBuffer([]byte(newWord)))
	w.Require().NoError(err)

	recorder := httptest.NewRecorder()

	//when
	w.serviceMock.EXPECT().SaveWord(gomock.Any()).Times(1).Return(expectedError)
	w.router.ServeHTTP(recorder, request)

	//then
	assert.Equal(w.T(), recorder.Code, http.StatusBadRequest)
	assert.Contains(w.T(), recorder.Body.String(), expectedError.Error())
}

func (w *WordTestSuite) TestAddWord_jsonMalformed_Failed() {
	//given
	newWord := "{\"Label\":\"hello\",\"Meaning\":\"ciao\",\"Sentence\":}"
	expectedErrorMessage := "body request malformed"

	request, err := http.NewRequest(http.MethodPost, "/word", bytes.NewBuffer([]byte(newWord)))
	w.Require().NoError(err)

	recorder := httptest.NewRecorder()

	//when
	w.serviceMock.EXPECT().SaveWord(gomock.Any()).Times(0)
	w.router.ServeHTTP(recorder, request)

	//then
	assert.Equal(w.T(), recorder.Code, http.StatusBadRequest)
	assert.Contains(w.T(), recorder.Body.String(), expectedErrorMessage)
}

func (w *WordTestSuite) TestListWords() {
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

	request, err := http.NewRequest(http.MethodGet, "/word", nil)
	w.Require().NoError(err)

	recorder := httptest.NewRecorder()

	//when
	w.serviceMock.EXPECT().ListWords().Times(1).Return(words, nil)
	w.router.ServeHTTP(recorder, request)

	//then
	assert.Equal(w.T(), http.StatusOK, recorder.Code)
	assert.Equal(w.T(), expectedWordsList, recorder.Body.String())
}

func (w *WordTestSuite) TestListWords_noAvailableDictionary_returnBadRequest() {
	//given
	expectedError := errors.New("no dictionary available")

	request, err := http.NewRequest(http.MethodGet, "/word", nil)
	w.Require().NoError(err)

	recorder := httptest.NewRecorder()

	//when
	w.serviceMock.EXPECT().ListWords().Times(1).Return(nil, expectedError)
	w.router.ServeHTTP(recorder, request)

	//then
	assert.Equal(w.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(w.T(), recorder.Body.String(), expectedError.Error())
}
