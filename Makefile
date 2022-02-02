GOFILES= $$(go list -f '{{join .GoFiles " "}}')

run:
	go run $(GOFILES)

build:
	go build -o $(GOPATH)/bin/realty-tool $(GOFILES)

setup:
	go run main.go setup=up

ratel:
	docker run -p 8000:8000 dgraph/ratel
