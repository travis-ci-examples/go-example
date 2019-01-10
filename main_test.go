package main

import "testing"

func TestSomethingPlease(t *testing.T) {
	if 1 == 0 {
		t.Fatal("oh no")
	}
}
