help:
	@cat Makefile | grep '^\w'

build:
	go build -o bin/m3utool main.go

