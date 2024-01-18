run: build
	./.bin/gistie

build:
	go build -o ./.bin/gistie ./src/cmd/main.go
