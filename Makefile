all: build

build:
	docker build -t convox/release .

release: build
	docker run -it --env-file=.env -v /var/run/docker.sock:/var/run/docker.sock convox/release
