package webui

import (
	"net/http"
	"strings"
	"time"
)

// LaunchAndServeString launches a Web based UI, and sets it up to display the value of ‘s’.
//
// This func is probably more often useful for those who want to test out webui, or debug.
func LaunchAndServeString(s string) error {
	serveHttp := func(w http.ResponseWriter, r *http.Request) {
		reader := strings.NewReader(s)

		http.ServeContent(w, r, "index.html", time.Now(), reader)
	}

	return LaunchAndServe( http.HandlerFunc(serveHttp) )
}
