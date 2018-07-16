package dispatch

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/larse514/dispatcher/handler"
)

//HTTPDispatcher is an implementation of the Dispatcher interface
type HTTPDispatcher struct {
	Client *http.Client
}

//DispatchMessage is a method to dispatch a message to a consumer
func (dispatcher HTTPDispatcher) DispatchMessage(message *handler.Message, route Route) error {

	req, err := http.NewRequest("POST", route.URL, bytes.NewBuffer([]byte(message.Message)))

	req.Header.Set("Content-Type", "application/json")

	resp, err := dispatcher.Client.Do(req)

	if err != nil {
		log.Println("ERROR: error making http call ", err)
		return errors.New("error making http call")
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)

	if is2xx(&resp.StatusCode) {
		return nil
	}

	log.Println("ERROR: returned response ", resp)
	return errors.New("Client returned http status code")
}

func is2xx(status *int) bool {
	switch *status {
	case 200:
		return true
	case 201:
		return true
	case 202:
		return true
	default:
		return false
	}
}
