package controller

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	mock_model "github.com/motnip/sermo/mocks/model"
	"github.com/motnip/sermo/model"
)

type DictionaryTestSuite struct {
	suite.Suite
	controller     *gomock.Controller
	repositoryMock *mock_model.MockRepository
	router         *mux.Router
	sut            DictionaryController
}

func TestDictionaryTestSuite(t *testing.T) {
	suite.Run(t, new(DictionaryTestSuite))
}

func (d *DictionaryTestSuite) SetupTest() {
	d.controller = gomock.NewController(d.T())
	d.repositoryMock = mock_model.NewMockRepository(d.controller)
	d.sut = NewController(d.repositoryMock)
	d.router = mux.NewRouter().StrictSlash(true)
	d.router.HandleFunc(d.sut.GetCreateDictionaryRoute().Path, d.sut.GetCreateDictionaryRoute().Function).Methods(d.sut.GetCreateDictionaryRoute().Method)
	d.router.HandleFunc(d.sut.GetListAllDictionary().Path, d.sut.GetListAllDictionary().Function).Methods(d.sut.GetListAllDictionary().Method)
}

func (d *DictionaryTestSuite) TestCreateDictionary() {
	//given
	dictionaryLanguage := "language"

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	d.Require().NoError(err)

	//when
	d.repositoryMock.EXPECT().CreateDictionary(gomock.Eq(dictionaryLanguage)).Times(1)
	d.router.ServeHTTP(recorder, request)

	assert.Equal(d.T(), http.StatusCreated, recorder.Code)
	assert.Equal(d.T(), dictionaryLanguage, recorder.Body.String())
}

func (d *DictionaryTestSuite) TestCreateDictionary_EmptyLanguage_returnBadRequest() {
	//given
	emptyDictionaryLanguage := ""
	expectedErrorMsg := "not valid language\n"

	request, err := http.NewRequest(http.MethodPost, "/dictionary", bytes.NewBuffer([]byte(emptyDictionaryLanguage)))
	d.Require().NoError(err)

	recorder := httptest.NewRecorder()

	//when
	d.repositoryMock.EXPECT().CreateDictionary(gomock.Any()).Times(0)
	d.router.ServeHTTP(recorder, request)

	//then
	assert.Equal(d.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(d.T(), expectedErrorMsg, recorder.Body.String())
}

func (d *DictionaryTestSuite) TestCreateDictionary_existDictionaryForALanguage_returnBadRequest() {
	//given
	dictionaryLanguage := "en"


	firstRequest, err := http.NewRequest(http.MethodPost, "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	d.Require().NoError(err)

	secondRequest, err := http.NewRequest(http.MethodPost, "/dictionary", bytes.NewBuffer([]byte(dictionaryLanguage)))
	d.Require().NoError(err)

	recorderFirstRequest := httptest.NewRecorder()
	recorderSecondRequest := httptest.NewRecorder()

	//when
	d.repositoryMock.EXPECT().CreateDictionary(gomock.Any()).Return(nil, nil)
	d.repositoryMock.EXPECT().CreateDictionary(gomock.Any()).Return(nil, errors.New("forced error"))
	d.router.ServeHTTP(recorderFirstRequest, firstRequest)
	d.router.ServeHTTP(recorderSecondRequest, secondRequest)

	//then
	assert.Equal(d.T(), http.StatusCreated, recorderFirstRequest.Code)
	assert.Equal(d.T(), http.StatusBadRequest, recorderSecondRequest.Code)
}

func (d *DictionaryTestSuite) TestListDictionary_Succeed() {
	//given
		dictionaryList := make([]*model.Dictionary, 0)
	words := make([]*model.Word, 0)
	words = append(words, &model.Word{
		Label:    "foo",
		Meaning:  "bar",
		Sentence: "var",
	})
	expectedReturn := model.Dictionary{
		Language: "en",
		Words:    words,
	}

	dictionaryList = append(dictionaryList, &expectedReturn)

	expectedList := "[{\"Language\":\"en\",\"Words\":[{\"Label\":\"foo\",\"Meaning\":\"bar\",\"Sentence\":\"var\"}]}]"

	request, err := http.NewRequest(http.MethodGet, "/dictionary", bytes.NewBuffer([]byte(expectedList)))
	d.Require().NoError(err)

	recorder := httptest.NewRecorder()
	//when
	d.repositoryMock.EXPECT().ListDictionary().Return(dictionaryList)
	d.router.ServeHTTP(recorder, request)

	//then
	assert.Equal(d.T(), http.StatusOK, recorder.Code)
	assert.Contains(d.T(), recorder.Body.String(), expectedList)
}
