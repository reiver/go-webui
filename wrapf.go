package webui

import (
	"bytes"
	"fmt"
	"errors"
)

// wrapf wraps an error in another error that is "branded" to webui.
//
// For example, this:
//
//	var err errors = errors.New("some error happened")
//	
//	const command string = "PUNCH"
//	
//	e := wrapf(err, "could not %q the thing", command)
//
//	fmt.Println(e)
//
// Would outout:
//
//	webui: could not "PUNCH" the thing: some error happened
func wrapf(err error, format string, a ...interface{}) error {
	var buffer bytes.Buffer

	buffer.WriteString("webui: ")

	fmt.Fprintf(&buffer, format, a...)

	if nil != err {
		buffer.WriteString(": ")

		buffer.WriteString(err.Error())
	}

	return errors.New(buffer.String())
}
