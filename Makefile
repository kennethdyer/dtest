
B=go build
I=go install
C=./cmd/dtest

all: tidy lint fmt install

install:
	$(I) $(C)

tidy:
	go mod tidy

lint: vet

vet:
	go vet

check:
	staticcheck

golint:
	golint ./...

fmt:
	gofmt -w .
