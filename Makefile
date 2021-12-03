build:
	subo build . --native --no-bundle

build/docker:
	subo build . --no-bundle

run:
	go run main.go

run/wasmtime:
	go run main.go -tags wasmtime

.PHONY: build build/docker