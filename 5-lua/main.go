package main

import (
	_ "embed"

	lua "github.com/yuin/gopher-lua"
)

//go:embed script.lua
var script string

func main() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(script); err != nil {
		panic(err)
	}
}
