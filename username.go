package webui

import (
	"os"
//	"os/user"
)

// username tries to figure out the user's system level username, and returns it.
//
// If it cannot figure out what it is, it returns an empty string ("").
func username() string {

	if name := os.Getenv("USERNAME"); "" != name {
		return name
	}

	if name := os.Getenv("USER"); "" != name {
		return name
	}

// When I try to build using this the ‘go’ compiler returns:
//
//	# runtime/cgo
//	exec: "gcc": executable file not found in $PATH
//
// So commenting out for now:....
//
//	{
//		user, err := user.Current()
//		if nil == err && nil != user {
//			if name := user.Username; "" != name {
//				return name
//			}
//		}
//	}

	return ""
}
