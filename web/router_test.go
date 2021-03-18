package web

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	mock_controller "sermo/mocks/controller"
	"testing"
)

func TestNewRouter(t *testing.T) {

	//given
	ctr := gomock.NewController(t)
	restFullController := mock_controller.NewMockController(ctr)

	//when
	sut := NewRouter(restFullController)

	//than
	sut.Init()

	if sut.router.GetRoute("createDictionary") == nil {
		t.Errorf("No path returend: got %v want %v", sut.router.Path("/dictionary"), "not nil")
	}

	if sut.router.GetRoute("createDictionary") == nil {
		t.Errorf("No path returend: got %v want %v", sut.router.Path("/dictionary"), "not nil")
	}
}

func TestRouter_InitRoute(t *testing.T) {

	//given
	ctr := gomock.NewController(t)
	mockController := mock_controller.NewMockController(ctr)

	newRoute := &Route{
		Path:     "/foo",
		Function: mockController.CreateDictionary,
		Method:   http.MethodGet,
		Name:     "fooGet",
	}

	expected := mux.NewRouter().StrictSlash(true)
	expected.HandleFunc(newRoute.Path, newRoute.Function).Name(newRoute.Name).Methods(newRoute.Method)

	//when
	sut := NewRouter(mockController)

	//than
	sut.InitRoute(newRoute)

	if sut.router.GetRoute("fooGet") == nil {
		t.Errorf("No path returend: got %v want %v", nil, "not nil")
	}

	if path, err := sut.router.GetRoute("fooGet").GetPathTemplate(); err == nil && path != "/foo" {
		t.Errorf("No path returend: got %v want %v", path, "/foo")
	}

	if sut.router.GetRoute("fooGet").GetHandler() == nil {
		t.Errorf("No path returend: got %v want %v", nil, "not nil")
	}
}
