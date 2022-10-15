package midi_parser

import (
	"fmt"
	"testing"
)

func expectedAndGotBool(t *testing.T, expected, got bool) {
	if got != expected {
		fmt.Println("FAIL! expected: ", expected, ", got: ", got)
		t.Fail()
	}
}

func TestIsRunningStatusLenOk(t *testing.T) {
	expectedAndGotBool(t, true, isRunningStatusLenOk(3, 2))
	expectedAndGotBool(t, false, isRunningStatusLenOk(3, 3))
	expectedAndGotBool(t, true, isRunningStatusLenOk(5, 2))
	expectedAndGotBool(t, false, isRunningStatusLenOk(6, 2))
	expectedAndGotBool(t, true, isRunningStatusLenOk(7, 2))
	expectedAndGotBool(t, true, isRunningStatusLenOk(1, 0))
	expectedAndGotBool(t, true, isRunningStatusLenOk(5, 0))
}
