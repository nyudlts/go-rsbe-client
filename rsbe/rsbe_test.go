package rsbe

import (
	"testing"
)

func TestRSBEVersion(t *testing.T) {
	want := "0.4.0"
	if Version != want {
		t.Errorf("Problem with version. Wanted: %s Got: %s", want, Version)
	}
}
