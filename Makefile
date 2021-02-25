help:
	@cat Makefile | grep '^\w'

run:
	go run main.go

build:
	go build -o bin/m3utool main.go

cross-build:
	for os in linux; do \
	for arch in amd64 arm; do \
	GOOS=$$os GOARCH=$$arch go build -o bin/m3utool_$$os_$$arch main.go; \
	done; \
	done

