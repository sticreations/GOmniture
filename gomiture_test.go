package GOmniture

import "testing"

func TestNew(t *testing.T) {
	omni := New("user", "secret")
	if omni.username != "user" || omni.sharedSecred != "secret" {
		t.Error("Authentication Header was not set")
	}
}
