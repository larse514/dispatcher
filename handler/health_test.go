package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingStatusOk(t *testing.T) {
	r := getRouter()
	r.GET("/ping", Ping)

	req, _ := http.NewRequest("GET", "/ping", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		_, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t.Log("Error parsing body")
			t.Fail()
		}
		return statusOK
	})

}
func TestPingResponseIsPong(t *testing.T) {
	r := getRouter()
	r.GET("/ping", Ping)

	req, _ := http.NewRequest("GET", "/ping", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {

		body, err := ioutil.ReadAll(w.Body)
		expected := `{"message":"pong"}`
		actual := string(body)
		pageOK := err == nil && actual == expected

		if err != nil {
			t.Fail()
		}
		return pageOK
	})

}
