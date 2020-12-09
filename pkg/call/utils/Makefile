install-go-fuzz:
	@go get -v -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build
go-fuzz-build:
	@mkdir -p ./fuzz/${func}
	@go-fuzz-build -func ${func} -o ./fuzz/${func}/harness.zip
go-fuzz-run:
	@go-fuzz -bin ./fuzz/${func}/harness.zip -workdir ./fuzz/${func} | tee ./fuzz/${func}/log.txt
