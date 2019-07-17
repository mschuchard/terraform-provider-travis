GOPATH := $(shell pwd)

all: dep clean build test install

dep:
	ln -sf $(shell pwd)/travis src/.
	GOPATH=$(GOPATH) go get

clean:
	rm -f terraform-provider-travis
	rm -rf src

build:
	GOPATH=$(GOPATH) go build -o test/terraform-provider-travis

test:
	cd test
	terraform init
	terraform plan

install:
	GOPATH=$(GOPATH) go install
