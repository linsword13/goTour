srcDir := ./goSquare
target := ./build/goSquare.dylib

all: build

.PHONY: all build

build:
	go build -o $(target) -buildmode=c-shared $(srcDir)

clean:
	rm $(target)