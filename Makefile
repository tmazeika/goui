main.wasm: wasm_exec.js
	GOOS=js GOARCH=wasm go build -o main.wasm

root := $(shell go env GOROOT)

wasm_exec.js:
	cp $(root)/misc/wasm/wasm_exec.js .
