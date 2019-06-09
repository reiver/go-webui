package webui

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

// LaunchAndServe launches a Web based UI, and sets it up to display what is coming out of ‘handler’.
func LaunchAndServe(handler http.Handler) error {

	listener, err := net.Listen("tcp", ":0")
	if nil != err {
		return wrapf(err, "problem getting next available TCP port for HTTP server")
	}

//@TODO: This might be using an external IP address, rather than the loop back (127.0.0.1).
//       Could this be an issue in any way?
	var addr string
	{
		netAddr := listener.Addr()
		if nil == netAddr {
			return wrapf(nil, "problem getting internal address structure which stores the TCP port for the HTTP server")
		}

		addr = listener.Addr().String()
	}

	// The ‘addr’ is going to contain something such as:
	//
	//	127.0.0.1:54239
	//
	// Or:
	//
	//	123.234.8.40:6752
	//
	// We we create an HTTP URL from it.
	var uri string = fmt.Sprintf("http://%s/", addr)

	httpDoneCh := make(chan struct{})
	httpErrCh  := make(chan error)

	go func(){

		if err := http.Serve(listener, handler); nil != err {
			httpErrCh <- err
			return
		}

		httpDoneCh <- struct{}{}
	}()

	launchDoneCh := make(chan struct{})
	launchErrCh  := make(chan error)

	go func(){
//@TODO: don't do it this way.
		time.Sleep(200 * time.Millisecond)

		if err := Launch(uri); nil != err {
			launchErrCh <- err
		}

		launchDoneCh <- struct{}{}
	}()

	select {
	case err := <-httpErrCh:
		return err
	case <- httpDoneCh:
		return nil
	case err := <-launchErrCh:
		return err
	case <- launchDoneCh:
		return nil
	}
}
