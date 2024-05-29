MODULE_NAME = report
CMD_DIR = ./src/cmd

.SILENT:
.DEFAULT_GOAL := help

.PHONY: help
help:
	$(info shylockgo commands:)
	$(info -> setup		installs dependencies)
	$(info -> format	formats go files)
	$(info -> build		builds executable binary)
	$(info -> run			runs application)

.PHONY: setup
setup:
	go get -d -v -t ./...
	go install -v ./...
	go mod tidy -v

.PHONY: format
format:
	go fmt ./...

.PHONY: build
build:
	go build -v -o $(MODULE_NAME) $(CMD_DIR)
	chmod +x $(MODULE_NAME)
	echo $(MODULE_NAME)

.PHONY: run
run:
	go run ./$(CMD_DIR) --all
