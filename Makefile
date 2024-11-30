.PHONY: all clean run

all: build

build:
	go build -o sand-sim

clean:
	rm -f sand-sim

run: build
	./sand-sim
