
B=go build
I=go install
C=./cmd/dtest

all: lint fmt install

install:
	$(I) $(C)

lint:
	golint $(C)

fmt:
	gofmt -w .
