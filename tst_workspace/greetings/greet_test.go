package greetings

import "testing"

func TestFailGreet(t *testing.T) {
	want := "Hello, world."
	if got := greet(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
