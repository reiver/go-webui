package webui

import (
	"bytes"
	"net/http"
	"time"
)

// LaunchAndServeBytes launches a Web based UI, and sets it up to display the value of ‘b’.
//
// This func is probably more often useful for those who want to test out webui, or debug.
func LaunchAndServeBytes(b []byte) error {
	serveHttp := func(w http.ResponseWriter, r *http.Request) {
		reader := bytes.NewReader(b)

		http.ServeContent(w, r, "index.html", time.Now(), reader)
	}

	return LaunchAndServe( http.HandlerFunc(serveHttp) )
}
