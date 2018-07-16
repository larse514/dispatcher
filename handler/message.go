package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	urlParam = "name"
)

//MessageHandler interface defines methods to integrate with event consumers
type MessageHandler interface {
	DispatchMessage(c *gin.Context)
}

//LambdaMessageHandler implementation of MessageHandler interface
type LambdaMessageHandler struct {
	Dispatcher Dispatcher
}

//Dispatcher interface to dispatch message to source
type Dispatcher interface {
	Dispatch(message *Message, source string) error
}

//Message to dispatch to consumers
type Message struct {
	Message string `json:"message"`
}

//DispatchMessage is method to handle dispatching message to consumers
func (handler LambdaMessageHandler) DispatchMessage(c *gin.Context) {
	message := Message{}
	sourceName := c.Params.ByName(urlParam)

	err := c.ShouldBind(&message)

	if err != nil {
		log.Println("ERROR: error parsing request ", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid request",
		})
		return
	}

	err = handler.Dispatcher.Dispatch(&message, sourceName)

	if err != nil {
		log.Println("ERROR: Error making http call with error ", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "Error making http call",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}
