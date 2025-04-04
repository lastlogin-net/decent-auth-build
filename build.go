package build

import (
	"io"
	"embed"
)

//go:embed decentauth.wasm
var fs embed.FS

var WasmBytes []byte

func init() {
	wasmFile, err := fs.Open("decentauth.wasm")
	if err != nil {
		panic(err)
	}

	WasmBytes, err = io.ReadAll(wasmFile)
	if err != nil {
		panic(err)
	}
}
