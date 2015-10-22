all: build

build:
	docker build -t convox/release .

cli:
	convox run --app release release cli

rack:
	convox run --app release release kernel

test:
	go test -v ./...
