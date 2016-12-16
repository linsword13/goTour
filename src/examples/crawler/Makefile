source := crawler.go

all: windows linux mac

.PHONY: all windows linux mac clean

deps:
	go get golang.org/x/net/html

clean:
	rm -rf build/

windows: deps
	GOOS=windows GOARCH=amd64 go build -o build/crawler_win_64.exe $(source)

linux: deps
	GOOS=linux GOARCH=amd64 go build -o build/crawler_linux_64 $(source)

mac: deps
	GOOS=darwin GOARCH=amd64 go build -o build/crawler_mac_64 $(source)