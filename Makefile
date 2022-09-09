.PHONY: default
default: build

all: clean get-deps test

version := "0.1.0"

test: 
	mkdir -p bin
	go test -coverprofile=coverage.out ./...
	# go tool cover -func=bin/cov.out
	go test -bench 'Benchmark' ./... | tee bin/bench.txt

clean:
	rm -rf ./bin
