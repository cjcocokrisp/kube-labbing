.PHONY: build

build-lambda-function:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(HANDLER) lambda/$(DIR)/$(DIR).go
	mv $(HANDLER) $(ARTIFACTS_DIR)

build-HelloWorldFunction: HANDLER=bootstrap
build-HelloWorldFunction: DIR=hello
build-HelloWorldFunction: build-lambda-function

build-FibFunction: HANDLER=bootstrap
build-FibFunction: DIR=fib
build-FibFunction: build-lambda-function

build-FacFunction: HANDLER=bootstrap
build-FacFunction: DIR=fac
build-FacFunction: build-lambda-function

build:
	sam build
