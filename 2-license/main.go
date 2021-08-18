package main

import (
	_ "embed"
	"fmt"
)

//go:embed MIT.txt
var license []byte

func main() {
	fmt.Println(string(license))
}
