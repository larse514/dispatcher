# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
CORE_BINARY_NAME=main
BINARY_NAME=main
SAM_OUTPUT=sam_output.yml
SAM_FILE=sam.yml
BINDATA_ASSETS=iaas.go

all: clean dependencies build test package

build: 
	# cd assets && ./go-bindata -prefix "ias/cloudformation/" -pkg assets -o $(BINDATA_ASSETS) ias/cloudformation/ 
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(CORE_BINARY_NAME) main.go
package:
	zip main.zip $(CORE_BINARY_NAME)
	
package-aws:
	aws cloudformation package --template-file $(SAM_FILE) --output-template-file $(SAM_OUTPUT) --s3-bucket $(DEPLOYMENT_BUCKET) 

test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME).zip
	rm -f $(SAM_OUTPUT)
deploy:
	aws cloudformation deploy --template-file $(SAM_OUTPUT) --stack-name $(STACK_NAME) --capabilities CAPABILITY_IAM --parameter-overrides Environment=$(ENV) TableBaseName=$(TABLE_NAME)

dependencies: 
	@go get github.com/aws/aws-lambda-go/lambda
	@go get github.com/aws/aws-lambda-go/events
	@go get github.com/gin-gonic/gin
	@go get github.com/awslabs/aws-lambda-go-api-proxy/gin
	@go get github.com/satori/go.uuid
	@go get github.com/aws/aws-sdk-go/aws
	@go get github.com/larse514/aws-cloudformation-go

integ: 
	cd integration && ./setup.sh $(STACK_NAME)
	cd integration && newman run Dispatcher_API.postman_collection.json