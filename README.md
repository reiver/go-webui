# go-webui

Package **webui** enables an application written in the Go programming language (i.e., Golang) to create a **user interface** (**UI**) using Web technologies, such HTML, CSS, JavaScript, WebAssembly, WebRTC, etc etc etc.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-webui

[![GoDoc](https://godoc.org/github.com/reiver/go-webui?status.svg)](https://godoc.org/github.com/reiver/go-webui)

## Hello World Example

A "hello world" example for using **webui** is:
```go
import "github.com/reiver/go-webui"

// ...

err := webui.LaunchAndServeString("Hello world!")
```

Or:
```go
import "github.com/reiver/go-webui"

// ...

err := webui.Launch("data:,Hello%20world!")
```

## More Realistic Example
A _more realistic_ example would probably have you (the programmer) create one or more `http.Handler`.

The `http.Handler` would serve the content that is shows in the Web UI.

Such as in:
```go
import "github.com/reiver/go-webui"

// ...

func serveHttp(w http.ResponseWriter, r *http.Request) {
	// ...
}

// ...

var handler http.Handler = http.HandlerFunc(serveHttp)

// ...

err := webui.LaunchAndServe(handler)
```

In other words, use can use many of the same tools you use to create Web applications, to create a desktop UI.
