.DEFAULT_GOAL := build

BINARY_NAME = svc
BUILD_PATH = cmd/build

build:
	mkdir -p $(BUILD_PATH)
	CGO_ENABLED=0 go build -o $(BUILD_PATH)/$(BINARY_NAME) main.go

clean:
	rm -rf $(BUILD_PATH)
