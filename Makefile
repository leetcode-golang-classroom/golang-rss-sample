.PHONY=build

build:
	@go build -o bin/main cmd/*.go

run: build
	@./bin/main

test:
	@go test -v -cover ./test/...