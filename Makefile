all:
	GOOS=js GOARCH=wasm go build
	GOOS=js GOARCH=wasm golangci-lint run --enable-all --exclude-use-default=false --disable=paralleltest