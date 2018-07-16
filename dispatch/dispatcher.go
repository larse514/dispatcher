package dispatch

import (
	"errors"
	"log"

	"github.com/larse514/dispatcher/handler"
)

//MessageDispatcher interface to dispatch message
type MessageDispatcher interface {
	DispatchMessage(message *handler.Message, route Route) error
}

//SourceClient interface to retrieve routes for a given source
type SourceClient interface {
	GetRoutes(sourcename string) ([]Route, error)
}

//LambdaDispatcher is an implementation of the MessageHandler interface
type LambdaDispatcher struct {
	SourceClient      SourceClient
	MessageDispatcher MessageDispatcher
}

//Route is a struct representing a Route
type Route struct {
	URL string `json:"url"`
}

//Dispatch is a method to dispatch a message to a source's routes
func (dispatcher LambdaDispatcher) Dispatch(message *handler.Message, source string) error {
	log.Println("DEBUG: about to retrieve routes from source ", source)
	routes, err := dispatcher.SourceClient.GetRoutes(source)

	if err != nil {
		log.Println("Error retrieving routes for a source ", err)
		return errors.New("Error fetching routes for source")
	}

	for _, route := range routes {
		log.Println("about to dispatch message ", message, " to ", route)
		err = dispatcher.MessageDispatcher.DispatchMessage(message, route)
	}

	return err
}
