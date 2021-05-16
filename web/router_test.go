package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	mock_controller "github.com/motnip/sermo/mocks/controller"
)

func TestRouter_InitRoute(t *testing.T) {

	//given
	ctr := gomock.NewController(t)
	mockController := mock_controller.NewMockDictionaryController(ctr)

	newRoute := &Route{
		Path:     "/foo",
		Function: mockController.CreateDictionary,
		Method:   http.MethodGet,
		Name:     "fooGet",
	}

	expected := mux.NewRouter().StrictSlash(true)
	expected.HandleFunc(newRoute.Path, newRoute.Function).Name(newRoute.Name).Methods(newRoute.Method)

	//when
	sut := NewRouter()

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

func TestRouter_requestWithWrongHeader_Failed(t *testing.T) {

	//given
	var headers = make(map[string]string)
	headers["Content-type"] = "application/json"
	newRoute := &Route{
		Path: "/foo",
		Function: func(rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(http.StatusOK)
		},
		Method:  http.MethodGet,
		Name:    "fooGet",
		Headers: &headers,
	}

	//when
	sut := NewRouter()

	//than
	sut.InitRoute(newRoute)

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, newRoute.Path, nil)
	request.Header.Add("Content-type", "text/plain")
	if err != nil {
		t.Fatal(err)
	}

	//when
	sut.router.ServeHTTP(recorder, request)

	//then
	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("Router returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
