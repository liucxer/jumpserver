all: jumpserver_mac jumpserver_linux jumpserver_windows

jumpserver_mac:
	mkdir -p build
	go build -o ./build/jumpserver_mac

jumpserver_linux:
	mkdir -p build && cd build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/jumpserver_linux

jumpserver_windows:
	mkdir -p build && cd build
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./build/jumpserver_windows

