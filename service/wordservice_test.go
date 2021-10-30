package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	mock_model "github.com/motnip/sermo/mocks/model"
	"github.com/motnip/sermo/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type WordServiceTestSuite struct {
	suite.Suite

	controller     *gomock.Controller
	repositoryMock *mock_model.MockRepository
	sut            WordService
	word           *model.Word
}

func TestWordServiceTestSuite(t *testing.T) {
	suite.Run(t, new(WordServiceTestSuite))
}

func (w *WordServiceTestSuite) SetupTest() {

	w.controller = gomock.NewController(w.T())
	w.repositoryMock = mock_model.NewMockRepository(w.controller)
	w.sut = NewWordService(w.repositoryMock)
	w.word = &model.Word{
		Label:    "label",
		Meaning:  "Meaning",
		Sentence: "sentence",
	}
}

func (w *WordServiceTestSuite) Test_saveWord_successful() {

	w.repositoryMock.EXPECT().ExistsDictionary().Times(1).Return(true)
	w.repositoryMock.EXPECT().AddWord(gomock.Any()).Return(nil)

	err := w.sut.SaveWord(w.word)

	assert.NoError(w.T(), err)
}

func (w *WordServiceTestSuite) Test_saveWord_errorOnSave_Failed() {

	expectedError := errors.New("forced error - save word")

	w.repositoryMock.EXPECT().ExistsDictionary().Times(1).Return(true)
	w.repositoryMock.EXPECT().AddWord(gomock.Any()).Times(1).Return(expectedError)

	err := w.sut.SaveWord(w.word)

	assert.Error(w.T(), err)
	assert.ErrorIs(w.T(), err, expectedError)
}

func (w *WordServiceTestSuite) Test_SaveWord_noDictionaryAvailable_Failed() {

	expectedError := model.NO_DICTIONARY

	w.repositoryMock.EXPECT().ExistsDictionary().Times(1).Return(false)
	w.repositoryMock.EXPECT().AddWord(gomock.Any()).Times(0)

	err := w.sut.SaveWord(w.word)
	assert.Error(w.T(), err)
	assert.EqualError(w.T(), err, expectedError)
}

func (w *WordServiceTestSuite) Test_saveWord_errorOnValidate_Failed() {

	w.word.Label = ""
	expectedError := errors.New("forced error - validation failed")

	w.repositoryMock.EXPECT().AddWord(gomock.Any()).Times(0).Return(expectedError)

	err := w.sut.SaveWord(w.word)

	assert.Error(w.T(), err, expectedError)
}

func (w *WordServiceTestSuite) Test_ListWords_succeed() {

	words := make([]*model.Word, 0)
	words = append(words, w.word)

	w.repositoryMock.EXPECT().ExistsDictionary().Times(1).Return(true)
	w.repositoryMock.EXPECT().ListWords().Times(1).Return(words, nil)

	savedWord, err := w.sut.ListWords()

	assert.NoError(w.T(), err)
	assert.Equal(w.T(), len(savedWord), len(words))
}

func (w *WordServiceTestSuite) Test_ListWords_noDictionaryAvailable_failed() {

	w.repositoryMock.EXPECT().ExistsDictionary().Times(1).Return(false)
	w.repositoryMock.EXPECT().ListWords().Times(0)

	savedWord, err := w.sut.ListWords()

	assert.Error(w.T(), err)
	assert.Errorf(w.T(), err, model.NO_DICTIONARY)
	assert.Len(w.T(), savedWord, 0)
}
