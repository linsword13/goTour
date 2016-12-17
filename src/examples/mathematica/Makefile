srcDir := ./goSquare
target := libgoSquare.so
target2 := libgoSquare.h

all: build

.PHONY: all build

build:
	go build -o $(target) -buildmode=c-shared $(srcDir)

clean:
	rm $(target) $(target2)