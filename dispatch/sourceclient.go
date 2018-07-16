package dispatch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//AWSSource is an implemenation of MessageDispatcher
type AWSSource struct {
	Client *http.Client
	URL    string
}

//SourceDTO response
type SourceDTO struct {
	Routes []Route `json:"routes"`
}

//GetRoutes is a method to retrieve routes for a source
func (source AWSSource) GetRoutes(sourcename string) ([]Route, error) {
	req, err := http.NewRequest("GET", source.URL, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := source.Client.Do(req)

	if err != nil {
		log.Println("ERROR: error making http call ", err)
		return nil, errors.New("error making http call")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	s := SourceDTO{}

	err = json.Unmarshal([]byte(body), &s)

	if err != nil {
		log.Println("ERROR: error unmarshalling response ", err)
		return nil, errors.New("Unmarshal error")
	}

	fmt.Println("response Status:", resp.Status)
	return s.Routes, nil
}
