test:
	go test -v ./...

lint: setup
	go vet ./...
	golint -set_exit_status ./...

.PHONY: setup test lint
