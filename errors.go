package webui

import (
	"errors"
)

var (
	errInternalError      = errors.New("webui: Internal Error")
	errNotFound           = errors.New("webui: Not Found")
	errWebBasedUINotFound = errors.New("webui: Web Based UI Not Found")
)
