GOOS		:= linux
GOARCH		:= amd64

.PHONY:		clean

all:		test

# .go files are reformatted to conform to gofmt standards
fmt:
		GOOS=$(GOOS) GOARCH=$(GOARCH) gofmt -d -e -s -w *.go

lint:		fmt
		GOOS=$(GOOS) GOARCH=$(GOARCH) golint -set_exit_status *.go

vet:		lint
		GOOS=$(GOOS) GOARCH=$(GOARCH) go tool vet *.go

test:		vet
		GOOS=$(GOOS) GOARCH=$(GOARCH) go test -race -coverprofile=coverage.txt -covermode=atomic -v .
		GOOS=$(GOOS) GOARCH=$(GOARCH) go tool cover -html=coverage.txt -o coverage.html

clean:
		@rm -f coverage.html
		@rm -f coverage.txt
