package web

import (
	"net/http"
	"net/http/httptest"
	"sermo/controller"
	"testing"
)

func TestAddPath(t *testing.T) {

	router := NewRouter()

	greetings := controller.NewGreetings()
	router.AddPath(&Route{
		Path:           greetings.Path(),
		ControllerFunc: greetings.Greetings,
	})

	if len(router.paths) < 1 {
		t.Error("no paths have not been persisted")
	}
}

func TestAddPath_addMultiplePath_succeed(t *testing.T) {

	router := NewRouter()

	greetings := controller.NewGreetings()
	router.AddPath(&Route{
		Path:           greetings.Path(),
		ControllerFunc: greetings.Greetings,
	})

	goodBy := controller.NewGoodbye()
	router.AddPath(&Route{
		Path:           goodBy.Path(),
		ControllerFunc: goodBy.Goodbye,
	})

	if len(router.paths) != 2 {
		t.Error("no paths have not been persisted")
	}

	if router.paths[0].Path != greetings.Path() {
		t.Errorf("expected %v, got %v", greetings.Path(), router.paths[0].Path)
	}

	if router.paths[1].Path != goodBy.Path() {
		t.Errorf("expected %v, got %v", goodBy.Path(), router.paths[1].Path)
	}
}

func TestAddPath_SingleCallPath_succeed(t *testing.T) {

	//given
	greetings := controller.NewGreetings()
	router := NewRouter()

	router.AddPath(&Route{
		Path:           greetings.Path(),
		ControllerFunc: greetings.Greetings,
	})
	router.RouterStart()

	req, err := http.NewRequest("GET", greetings.Path(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(greetings.Greetings).ServeHTTP(rr, req)
	//handler.ServeHTTP()

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "Welcome home!"
	responseBody := rr.Body.String()
	if responseBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, expected)
	}
}

func Test_MultipleCallPath_succeed(t *testing.T) {

	//given
	router := NewRouter()

	greetings := controller.NewGreetings()

	goodBye := controller.NewGoodbye()
	router.AddPath(&Route{
		Path:           goodBye.Path(),
		ControllerFunc: goodBye.Goodbye,
	})

	router.AddPath(&Route{
		Path:           greetings.Path(),
		ControllerFunc: greetings.Greetings,
	})

	//	router.RouterStart()

	// greeting test
	req, err := http.NewRequest("GET", greetings.Path(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(greetings.Greetings).ServeHTTP(rr, req)
	//handler.ServeHTTP()

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "Welcome home!"
	responseBody := rr.Body.String()
	if responseBody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBody, expected)
	}

	//goodbye test
	reqGoodBy, err := http.NewRequest("GET", "/stocazzzo", nil)
	if err != nil {
		t.Fatal(err)
	}

	rrGoodBy := httptest.NewRecorder()
	http.HandlerFunc(goodBye.Goodbye).ServeHTTP(rrGoodBy, reqGoodBy)
	//handler.ServeHTTP()

	if status := rrGoodBy.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expectedGoodBy := "Good Bye!"
	responseBodyGoodBy := rrGoodBy.Body.String()
	if responseBodyGoodBy != expectedGoodBy {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseBodyGoodBy, expectedGoodBy)
	}
}
