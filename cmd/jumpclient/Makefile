all: jumpclient_mac jumpclient_linux jumpclient_windows

jumpclient_mac:
	mkdir -p build
	go build -o ./build/jumpclient_mac

jumpclient_linux:
	mkdir -p build && cd build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/jumpclient_linux

jumpclient_windows:
	mkdir -p build && cd build
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./build/jumpclient_windows

