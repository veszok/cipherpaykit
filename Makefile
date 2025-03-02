.DEFAULT_GOAL := build
.PHONY:fmt vet build test
fmt: 
	go fmt ./...  
vet: fmt 
	go vet ./...  
build: vet 
	go build
test: fmt vet
	go test ./...
coverage: build
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out