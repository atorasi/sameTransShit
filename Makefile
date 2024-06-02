.PHONY: run
run:
	go run ./cmd/sametrans

.PHONY: build
build: 
	go build -o ./build/SameTrans -v ./cmd/sametrans

.SILENT:
.DEFAULT_GOAL := run
