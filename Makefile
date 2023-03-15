#
# Variables
#

DOCKER_NAME=dstroot/collection_machine
VERSION=1.0.5
SHELL=/bin/bash

#
# Build
#

all: $(info Current version is $(VERSION)) build version push clean

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags "-X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash=`git rev-parse HEAD` -w -s" .
	docker build -t $(DOCKER_NAME):latest .

version:
	docker tag $(DOCKER_NAME):latest $(DOCKER_NAME):$(VERSION)

push:
	docker push $(DOCKER_NAME):latest
	docker push $(DOCKER_NAME):$(VERSION)

clean:
	rm -rf $(DOCKER_NAME) && rm -rf $(DOCKER_NAME).exe
