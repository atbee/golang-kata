package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, atbee!"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
