package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

const port = "8888"

//go:embed static
var s embed.FS

func main() {
	fsys, _ := fs.Sub(s, "static")
	http.Handle("/", http.FileServer(http.FS(fsys)))
	fmt.Printf("Point your browser to http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
