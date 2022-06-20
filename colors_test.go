package colors

import "testing"

func TestColorize(t *testing.T) {
	got := Colorize("test", FgGreen)
	want := "\x1b[32mtest\x1b[0m"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = Colorize("test", FgGreen, BgRed)
	want = "\x1b[32;41mtest\x1b[0m"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
