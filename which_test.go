package webui

import (
	"fmt"
	"os"
	"strings"

	"testing"
)

func TestWhichErrNotFound(t *testing.T) {

	tests := []struct{
		Name string
	}{
		{
			Name: "qwertyuiopasdfghjklzxcvbnm",
		},
		{
			Name: "zxzxzxzxzxzxzxzaac",
		},
		{
			Name: "oknsdfnkwef",
		},
		{
			Name: "fib234897ydfshi23r98",
		},
		{
			Name: "327ym-983crfd328r",
		},
		{
			Name: "122121-sdss-fdgeb",
		},
	}

	for testNumber, test := range tests {

		actual, err := which(test.Name)
		if nil == err {
			t.Errorf("For test #%d (name =%q), expected an error, but did not actually got one: %#v", testNumber, test.Name, err)
			continue
		}
		if expected, actual := errNotFound, err; expected != actual {
			t.Errorf("For test #%d (name =%q), expected error (%T) %q, but actually got (%T) %q", testNumber, test.Name, expected, expected, actual, actual)
			continue
		}

		if expected := ""; expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q", testNumber, expected, actual)
			continue
		}
	}
}

// =========================================================================== //
//                                                                             //
// NOTE: This test should only really be run on Linux, and Unix-like systems.  //
//                                                                             //
// =========================================================================== //
func TestWhich(t *testing.T) {

	tests := []struct{
		Name string
	}{
		{
			Name: "clear",
		},
		{
			Name: "locate",
		},
		{
			Name: "ls",
		},
		{
			Name: "man",
		},
		{
			Name: "mkdir",
		},
		{
			Name: "mv",
		},
		{
			Name: "rm",
		},
		{
			Name: "rmdir",
		},
		{
			Name: "touch",
		},
	}

	for testNumber, test := range tests {

		actual, err := which(test.Name)
		if nil != err {
			t.Errorf("For test #%d (name =%q), did not expect an error, but actually got one: (%T) %q", testNumber, test.Name, err, err)
			continue
		}

		if expectedEndsWith, actual := fmt.Sprintf("%s%s", string(os.PathSeparator), test.Name), actual; !strings.HasSuffix(actual, expectedEndsWith) {
			t.Errorf("For test #%d, expected %q to end with %q, but didn't.", testNumber, actual, expectedEndsWith)
			continue
		}
	}
}
