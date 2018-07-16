package dispatch

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
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

	url := strings.Replace(source.URL, ":name", sourcename, -1)
	log.Println("DEBUG: created url ", url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := source.Client.Do(req)

	if err != nil {
		log.Println("ERROR: error making http call to source with error", err)
		return nil, errors.New("error making http call")
	}

	defer resp.Body.Close()
	s := SourceDTO{}
	err = json.NewDecoder(resp.Body).Decode(&s)

	if err != nil {
		log.Println("ERROR: error unmarshalling response ", err)
		return nil, errors.New("Unmarshal error")
	}

	fmt.Println("response Status:", resp.Status)
	log.Println("returning ", s.Routes)
	return s.Routes, nil
}
