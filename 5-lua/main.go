package main

import (
	_ "embed"

	lua "github.com/yuin/gopher-lua"
)

//go:embed script.lua
var script []byte

func main() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(string(script)); err != nil {
		panic(err)
	}
}
