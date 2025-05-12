
B=go build
I=go install
C=./cmd/dtest

all: tidy lint vet fmt install

install:
	$(I) $(C)

tidy:
	go mod tidy

lint: 
	golint ./...

vet:
	go vet ./...

check:
	staticcheck

golint:
	golint ./...

fmt:
	gofmt -w .
