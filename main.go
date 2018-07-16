package main

import (
	"net/http"
	"os"

	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/larse514/dispatcher/dispatch"
	"github.com/larse514/dispatcher/handler"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	sourceURLKey = "SourceUrl"
)

var ginLambda *ginadapter.GinLambda

// Handler is the main entry point for Lambda. Receives a proxy request and
// returns a proxy response
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r := gin.Default()
	sourceURL := os.Getenv(sourceURLKey)

	client := http.Client{}
	dispatchClient := dispatch.HTTPDispatcher{Client: &client}
	sourceClient := dispatch.AWSSource{Client: &client, URL: sourceURL}

	dispatcher := dispatch.LambdaDispatcher{SourceClient: sourceClient, MessageDispatcher: dispatchClient}
	h := handler.LambdaMessageHandler{Dispatcher: dispatcher}
	r.POST("/sources/:name/routes/messages", h.DispatchMessage)
	r.GET("/ping", handler.Ping)
	ginLambda = ginadapter.New(r)

	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
