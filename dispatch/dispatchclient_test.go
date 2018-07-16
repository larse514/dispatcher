package dispatch

import (
	"net/http"
)

type mock200HttpClient struct {
}

func (c *mock200HttpClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: 200}, nil
}

//TODO- add tests
// func TestNoErrorsReturnsNil(t *testing.T) {
// 	//arrange

// 	//act
// 	//assert
// }
