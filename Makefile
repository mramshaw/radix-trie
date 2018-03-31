GOPATH		:= /go
GOOS		:= linux
GOARCH		:= amd64

.PHONY:		run, clean

all:		test

# .go files are reformatted to conform to gofmt standards
fmt:
		GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) gofmt -d -e -s -w *.go

lint:		fmt
		GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) golint -set_exit_status *.go

vet:		lint
		GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) go tool vet *.go

test:		vet
		GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) go test -v
