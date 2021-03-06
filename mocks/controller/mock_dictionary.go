// Code generated by MockGen. DO NOT EDIT.
// Source: controller/dictionary.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDictionaryController is a mock of DictionaryController interface.
type MockDictionaryController struct {
	ctrl     *gomock.Controller
	recorder *MockDictionaryControllerMockRecorder
}

// MockDictionaryControllerMockRecorder is the mock recorder for MockDictionaryController.
type MockDictionaryControllerMockRecorder struct {
	mock *MockDictionaryController
}

// NewMockDictionaryController creates a new mock instance.
func NewMockDictionaryController(ctrl *gomock.Controller) *MockDictionaryController {
	mock := &MockDictionaryController{ctrl: ctrl}
	mock.recorder = &MockDictionaryControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDictionaryController) EXPECT() *MockDictionaryControllerMockRecorder {
	return m.recorder
}

// CreateDictionary mocks base method.
func (m *MockDictionaryController) CreateDictionary(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateDictionary", httpResponse, httpRequest)
}

// CreateDictionary indicates an expected call of CreateDictionary.
func (mr *MockDictionaryControllerMockRecorder) CreateDictionary(httpResponse, httpRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDictionary", reflect.TypeOf((*MockDictionaryController)(nil).CreateDictionary), httpResponse, httpRequest)
}
