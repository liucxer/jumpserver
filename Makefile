all:
	cd cmd/jumpcli && make && cd -
	cd cmd/jumpclient && make && cd -
	cd cmd/jumpserver && make && cd -

clean:
	rm -rf `find ./ -name "build"`