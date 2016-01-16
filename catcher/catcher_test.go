package main

import "testing"

func expect(t *testing.T, logText string, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("%v. Expected %v, got %v.", logText, expected, actual)
	}
}

func expectFatal(t *testing.T, logText string, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Fatalf("%v. Expected %v, got %v.", logText, expected, actual)
	}
}

func assert(t *testing.T, logText string, cond bool) {
	if !cond {
		t.Error(logText)
	}
}

func assertFatal(t *testing.T, logText string, cond bool) {
	if !cond {
		t.Fatal(logText)
	}
}
