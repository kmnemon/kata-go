package testframe

import (
	"testing"
)

func TestWasRun(t *testing.T) {
	var test *WasRun = NewWasRun((*WasRun).testMethod)
	if test.wasRun == true {
		t.Error("wasRun should be false")
	}

	test.run()

	if test.wasRun == false {
		t.Error("wasRun should be true")
	}
}
