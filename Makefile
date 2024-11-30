.PHONY: all clean run

all: build

build:
	go build -o ./bin/sand-sim

clean:
	rm -f ./bin/sand-sim

run: build
	./bin/sand-sim
