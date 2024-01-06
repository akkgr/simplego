PACKAGES := $(shell go list ./...)
name := $(shell basename ${PWD})

all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: start
start: css
		go run .

## dev: run dev server
.PHONY: dev
dev:
	go run .

## css: build tailwindcss
.PHONY: css
css:
	npx tailwindcss -i input.css -o static/css/main.css --minify

## css-watch: watch build tailwindcss
.PHONY: css-watch
css-watch:
	npx tailwindcss -i input.css -o static/css/main.css --watch