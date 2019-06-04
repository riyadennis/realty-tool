GOFILES= $$(go list -f '{{join .GoFiles " "}}')

run:
	go run $(GOFILES)

build:
	go build -o $(GOPATH)/bin/realty-tool $(GOFILES)

setup:
	go run main.go setup=up
