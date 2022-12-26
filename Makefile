.PHONY: build 
build: 
	go build -v ./cmd/apiserver
.PHONY: test
test: 
	go test -v ./internal/app/apiserver
	go test -v ./internal/app/apiserver/store
	go test -v ./internal/app/model
	


.DEFAULT_GOAL := build