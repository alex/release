all: build

build:
	docker build -t convox/release .

release:
	convox run release kernel
