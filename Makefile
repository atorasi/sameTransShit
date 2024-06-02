.PHONY: run
run:
	go run ./cmd/dna

.PHONY: build
build: 
	go build -o ./build/DNA -v ./cmd/dna

.SILENT:
.DEFAULT_GOAL := run
