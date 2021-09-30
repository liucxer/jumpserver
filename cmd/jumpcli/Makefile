all: jumpcli_mac jumpcli_linux jumpcli_windows

jumpcli_mac:
	mkdir -p build
	go build -o ./build/jumpcli_mac

jumpcli_linux:
	mkdir -p build && cd build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/jumpcli_linux

jumpcli_windows:
	mkdir -p build && cd build
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./build/jumpcli_windows

