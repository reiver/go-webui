package webui

import (
	"bytes"
	"io"
	"os/exec"
	"unicode/utf8"
)

// which calls the ‘which’ command, that is common on Linux, and Unix like systems.
//
//
// For example, calling:
//
//	which("nano")
//
// Is equivalent to calling:
//
//	which nano
//
// This will probably return the path:
//
//	/bin/nano
//
//
// For example, calling:
//
//	which("ssh")
//
// Is equivalent to calling:
//
//	which ssh
//
// This will probably return the path:
//
//	/usr/bin/ssh
//
//
// And calling:
//
//	which("mkdir")
//
// Is equivalent to calling:
//
//	which mkdir
//
// This will probably return the path:
//
//	/bin/mkdir
//
//
// And calling:
//
//	which("which")
//
// Is equivalent to calling:
//
//	which which
//
// This will probably return the path:
//
//	/usr/bin/which
//
//
// Etc.
func which(name string) (string, error) {

	const cmdName string = "which"
	cmd := exec.Command(cmdName, name)
	if nil == cmd {
		return "", errInternalError
	}

	var stdout io.ReadCloser
	{
		var err error

		stdout, err = cmd.StdoutPipe()
		if nil != err {
			return "", wrapf(err, "problem getting access to output from %q command", cmdName)
		}

		if err = cmd.Start(); nil != err {
			return "", wrapf(err, "problem starting %q command", cmdName)
		}
	}

	var buffer bytes.Buffer
	{
		_, err := buffer.ReadFrom(stdout)
		if nil != err {
			return buffer.String(), wrapf(err, "problem reading from output from %q command", cmdName)
		}
	}

	var err error
	{
		err = cmd.Wait()

		if nil != err {
			switch err.(type) {
			case *exec.ExitError:
				switch err.Error() {
				case "exit status 1":
					err = errNotFound
				case "exit status 2":
					err = errInternalError
				default:
					// Nothing here
				}
			}
		}
	}

	var p []byte = buffer.Bytes()
	func(){

		if 1 > len(p) {
			return
		}

		r, size := utf8.DecodeLastRune(p)
		if utf8.RuneError == r {
			return
		}


		if '\n' != r {
			return
		}

		p = p[:len(p)-size]
	}()

	return string(p), err
}
