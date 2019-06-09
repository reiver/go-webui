package webui

import (
	"errors"

	"testing"
)

func TestWrapf(t *testing.T) {

	tests := []struct{
		Err      error
		Format   string
		Args   []interface{}
		Expected string
	}{
		{
			Err: nil,
			Format: "problem!",
			Args: []interface{}(nil),
			Expected: "webui: problem!",
		},
		{
			Err: errors.New("pkg: oops, had a problem"),
			Format: "problem!",
			Args: []interface{}(nil),
			Expected: "webui: problem!: pkg: oops, had a problem",
		},
	}

	for testNumber, test := range tests {

		actual := wrapf(test.Err, test.Format, test.Args...)
		if expected, actual := test.Expected, actual.Error(); expected != actual {
			t.Errorf("For test #%d...", testNumber)
			t.Errorf("\tEXPECTED: %q", expected)
			t.Errorf("\tACTUAL:   %q", actual)
			continue
		}
	}
}
