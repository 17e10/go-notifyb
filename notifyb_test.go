package notifyb

import "testing"

func TestNotify(t *testing.T) {
	const e Notify = "my notify"
	got := e.Error()
	want := "my notify"
	if got != want {
		t.Errorf("Notify(%q) = %q, want %q", string(e), got, want)
	}
}
