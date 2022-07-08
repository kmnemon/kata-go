package tddselftest

import "testing"

func TestWasRun(t *testing.T) {
	var test WasRun
	if test.wasRun == true {
		t.Error("wasRun should be false")
	}

	test.run()

	if test.wasRun == false {
		t.Error("wasRun should be true")
	}
}
