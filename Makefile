#APP := $(shell basename $(shell git remote get-url origin))
#LOWER_APP := $(shell echo $(APP) | tr '[:upper:]' '[:lower:]')
REGISTRY := distefano119/tbot
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux #darwin windows
TARGETARCH=arm64 #amd64 if you have macOS: dpkg --print-architecture | cut -d'-' -f2

format:
		gofmt -s -w ./

lint:
		golint

test:
		go test -v

get:
		go get

build: format get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o tBot -ldflags "-X="github.com/distefano119ua/tBot/cmd.appVersion=${VERSION}

image:
		docker build . -t ${REGISTRY}:${VERSION}-${TARGETARCH}

push:
		docker push ${REGISTRY}:${VERSION}-${TARGETARCH}

clean:
		docker rmi ${REGISTRY}:${VERSION}-${TARGETARCH}