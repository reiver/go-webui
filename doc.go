/*
Package webui enables an application written in the Go programming language (i.e., Golang) to create a user interface (UI) using Web technologies, such HTML, CSS, JavaScript, WebAssembly, WebRTC, etc etc etc.

Hello World Example

A "hello world" example for using webui is:

	import "github.com/reiver/go-webui"
	
	// ...
	
	err := webui.LaunchAndServeString("Hello world!")

Or:

	import "github.com/reiver/go-webui"
	
	// ...
	
	err := webui.Launch("data:,Hello%20world!")

More Realistic Example

A more realistic example would probably have you (the programmer) create one or more ‘http.Handler’.

The ‘http.Handler’ would serve the content that is shows in the Web UI.

Such as in:

	import "github.com/reiver/go-webui"
	
	// ...
	
	func serveHttp(w http.ResponseWriter, r *http.Request) {
		// ...
	}
	
	// ...
	
	var handler http.Handler = http.HandlerFunc(serveHttp)
	
	// ...
	
	err := webui.LaunchAndServe(handler)

In other words, use can use many of the same tools you use to create Web applications, to create a desktop UI.
*/
package webui
