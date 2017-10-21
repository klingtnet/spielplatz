package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, World!"
	if msg() != expected {
		t.Fatalf("Expected %q but was %q", expected, msg())
	}
}
