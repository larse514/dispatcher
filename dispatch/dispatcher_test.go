package dispatch

import (
	"errors"
	"testing"

	"github.com/larse514/dispatcher/handler"
)

const (
	sourceClientError    = "SOURCECLIENTERROR"
	messageDispatchError = "MESSAGEDISPATCHERROR"
)

type goodSourceClient struct {
}
type mockMessageDispatcher struct {
}

func (client goodSourceClient) GetRoutes(source string) ([]Route, error) {
	return append(make([]Route, 0), Route{URL: "URL"}), nil
}

func (dispatcher mockMessageDispatcher) DispatchMessage(message *handler.Message, route Route) error {
	return nil
}

type mockBadSourceClient struct {
}
type mockBadMessageDispatcher struct {
}

func (client mockBadSourceClient) GetRoutes(source string) ([]Route, error) {
	return nil, errors.New(sourceClientError)
}

func (dispatcher mockBadMessageDispatcher) DispatchMessage(message *handler.Message, route Route) error {
	return errors.New(messageDispatchError)
}

func TestNoErrorsReturnsNil(t *testing.T) {
	//arrange
	dispatcher := LambdaDispatcher{SourceClient: goodSourceClient{}, MessageDispatcher: mockMessageDispatcher{}}
	//act
	err := dispatcher.Dispatch(&handler.Message{}, "source")
	//assert

	if err != nil {
		t.Log("Error received when non expected ", err)
		t.Fail()
	}
}

func TestSourceClientReturnsErrorReturnError(t *testing.T) {
	//arrange
	dispatcher := LambdaDispatcher{SourceClient: mockBadSourceClient{}, MessageDispatcher: mockMessageDispatcher{}}
	expected := "Error fetching routes for source"
	//act
	actual := dispatcher.Dispatch(&handler.Message{}, "source").Error()
	//assert

	if actual != expected {
		t.Log("Expected ", expected, " got ", actual)
		t.Fail()
	}
}

func TestMessageDispatcherReturnsErrorReturnError(t *testing.T) {
	//arrange
	dispatcher := LambdaDispatcher{SourceClient: goodSourceClient{}, MessageDispatcher: mockBadMessageDispatcher{}}
	expected := messageDispatchError
	//act
	actual := dispatcher.Dispatch(&handler.Message{}, "source").Error()
	//assert

	if actual != expected {
		t.Log("Expected ", expected, " got ", actual)
		t.Fail()
	}
}
