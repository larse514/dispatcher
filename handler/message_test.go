package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockMessageHandler struct {
	Dispatcher Dispatcher
}

func (mock MockMessageHandler) Dispatch(message *Message, source string) error {
	return nil
}

type MockBadMessageHandler struct {
	Dispatcher Dispatcher
}

func (mock MockBadMessageHandler) Dispatch(message *Message, source string) error {
	return errors.New("AN ERROR")
}

func TestDispatchMessageSuccessReturnsStatusOk(t *testing.T) {
	r := getRouter()

	handler := LambdaMessageHandler{Dispatcher: MockMessageHandler{}}
	r.POST("/sources/:name/routes/messages", handler.DispatchMessage)

	req, _ := http.NewRequest("POST", "/sources/sourcename/routes/messages",
		strings.NewReader(`{"message": "SAMPLEMESSAGE"}`))

	expected := http.StatusOK

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		actual := w.Code

		if actual != expected {
			t.Log("Expected ", expected, " but got ", actual)
			t.Fail()
		}
		return actual == expected
	})

}

func TestDispatchMessageFailsReturnServiceUnavailable(t *testing.T) {
	r := getRouter()

	handler := LambdaMessageHandler{Dispatcher: MockBadMessageHandler{}}
	r.POST("/sources/:name/routes/messages", handler.DispatchMessage)

	req, _ := http.NewRequest("POST", "/sources/sourcename/routes/messages",
		strings.NewReader(`{"message": "SAMPLEMESSAGE"}`))

	expected := http.StatusServiceUnavailable

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		actual := w.Code

		if actual != expected {
			t.Log("Expected ", expected, " but got ", actual)
			t.Fail()
		}
		return actual == expected
	})

}
