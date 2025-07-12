.PHONY: build
build:
	go build -o cmd/**/*.go

.PHONY: run
run:
	go run cmd/**/*.go
