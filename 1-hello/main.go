package main

import (
	_ "embed" // Import embed for its side effect (//go:embed directive)
	"fmt"
)

//go:embed hello.txt
var hello string

func main() {
	fmt.Println(hello)
}
