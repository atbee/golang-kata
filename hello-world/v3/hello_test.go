package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Foo")
	want := "Hello, Foo"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
