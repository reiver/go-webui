package webui

import (
	"fmt"
	"os/exec"
)

// Launch launches a Web based UI, specified by ‘uri’.
//
// For example:
//
//	err := webui.Launch("http://www.example.com/")
//
// Or, for example:
//
//	err := webui.Launch("http://127.0.0.1:8080/")
//
// Or also, for example:
//
//	err := webui.Launch("data:,Hello%20world!")
func Launch(uri string) error {

	if err := launchChromium(uri); nil == err {
		return nil
	}

//	if err := launchFirefox(uri); nil == err {
//		return nil
//	}

	return errWebBasedUINotFound
}

func launchChromium(uri string) error {

	path, err := discoverChromium()
	if nil != err {
		return err
	}

	// This is equivalent to:
	//
	// 	chromium --incognito --app=$uri
//@TODO: Do we need to escape ‘uri’?
	cmd := exec.Command(path, "--incognito", fmt.Sprintf("--app=%s",uri))
	if nil == cmd {
		return errInternalError
	}

	if err = cmd.Run(); nil != err {
		return wrapf(err, "problem running %q command", path)
	}

	return nil
}
