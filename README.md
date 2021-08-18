# Golang Embed

## What I like in Go

* Fast Compilation
* Simple Language (*)
* Cross compilation
* **Single Executable**

## Embed

Available since **Go 1.16**

From https://pkg.go.dev/embed#hdr-Directives:

> Package `embed` provides access to files embedded in the running Go program.

> Go source files that import `embed` can use the `//go:embed` directive to initialize a variable of type `string`, `[]byte`, or `FS` with the **contents of files** read from the package directory or subdirectories at **compile time**.

* Variable of type `string`
* Variable fo type `[]byte` for an individual file
* Variable of type `embed.FS` for an individual file or a group of files

## Examples

### 1 - Hello Example

* Variable of type `string`

```go
...

import (
	_ "embed" // Import embed for its side effect (//go:embed directive)
)

//go:embed hello.txt
var hello string

...
```

### 2 - License Example

* Variable of type `[]bytes`


```go
...

import (
	_ "embed"
)

//go:embed MIT.txt
var license []byte

...
```

### 3 - Config Example

* Variable of type `embed.FS` (with a single file)

```go
//go:embed config.json
var f embed.FS

func main() {
	data, _ := f.ReadFile("config.json")
...
```

### 4 - Static Web Site Example

* Variable of type `embed.FS` (with a subdirectory)

```go
...
//go:embed static
var static embed.FS

func main() {
    // Sub() returns a FS interface corresponding 
	// to the content of the `static` sub directory
	fsys, _ := fs.Sub(static, "static")

    // http.FS convert an FS implementation
	// Corresponds to http.Dir("/static") of operating file system
	http.Handle("/", http.FileServer(http.FS(fsys)))
	http.ListenAndServe(":8888", nil)
}
```

### 5 - Embedded Lua Example

* https://github.com/andreburgaud/meetup-golang-lua

```go
import (
	_ "embed"
	lua "github.com/yuin/gopher-lua"
)

//go:embed script.lua
var script []byte

func main() {
	L := lua.NewState()        // Opens Lua
	defer L.Close()
	L.DoString(string(script)) // Executes Lua string
}
```

## Resources

* https://pkg.go.dev/embed
* https://bhupesh-v.github.io/embedding-static-files-in-golang/
* https://blog.carlmjohnson.net/post/2021/how-to-use-go-embed/
* https://github.com/carlmjohnson/exembed/
* https://blog.jetbrains.com/go/2021/06/09/how-to-use-go-embed-in-go-1-16/
* https://lakefs.io/working-with-embed-in-go/
* https://echorand.me/posts/go-embed/