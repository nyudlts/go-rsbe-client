package rsbe

import (
	"strings"
	"testing"
)

// copied from Cobra test code:
// https://github.com/spf13/cobra/blob/40d34bca1bffe2f5e84b18d7fd94d5b3c02275a6/command_test.go#L49
func checkStringContains(t *testing.T, got, expected string) {
	if !strings.Contains(got, expected) {
		t.Errorf("Expected to contain: \n %v\nGot:\n %v\n", expected, got)
	}
}
