all: build

build:
	docker build -t convox/release .

test:
	go test -v ./...