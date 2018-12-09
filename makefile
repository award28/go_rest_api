# Go parameters
GO=go
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
GOGET=$(GO) get
APP_DIR=go_rest_api
PATH=$(GOPATH)/src/$(APP_DIR)
BINARY_NAME=go_rest_api
APP_PATH=$(PATH)/cmd/app/$(BINARY_NAME).go

all: install run
install: 
	$(GOINSTALL) $(APP_PATH)
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(GOBIN)/$(BINARY_NAME)
run:	
	$(GOBIN)/$(BINARY_NAME)
deps:
	$(GOGET) github.com/markbates/bolt
	$(GOGET) github.com/markbates/pop
