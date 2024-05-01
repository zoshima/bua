default: build run

build:
	go build -o bc

run:
	./bc --auth ""
