// File: testutils/assertions.go
package testutils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// AssertEquivalentTimestamps ensures two time strings represent the same moment.
func AssertEquivalentTimestamps(t *testing.T, expectedStr, actualStr string, msgAndArgs ...any) bool {
	t.Helper() // Crucial: tells Go this is a helper, so failures point to the actual test!

	expectedTime, err1 := time.Parse(time.RFC3339Nano, expectedStr)
	if !assert.NoError(t, err1, "Expected string is not a valid RFC3339 timestamp") {
		return false
	}

	actualTime, err2 := time.Parse(time.RFC3339Nano, actualStr)
	if !assert.NoError(t, err2, "Actual string is not a valid RFC3339 timestamp") {
		return false
	}

	// Compare them, allowing up to 1 millisecond of difference
	return assert.WithinDuration(t, expectedTime, actualTime, time.Millisecond, msgAndArgs...)
}

func AssertNotEquivalentTimestamps(t *testing.T, expectedStr, actualStr string, msgAndArgs ...any) bool {
	t.Helper()

	expectedTime, err1 := time.Parse(time.RFC3339Nano, expectedStr)
	if !assert.NoError(t, err1, "Expected string is not a valid RFC3339 timestamp") {
		return false
	}

	actualTime, err2 := time.Parse(time.RFC3339Nano, actualStr)
	if !assert.NoError(t, err2, "Actual string is not a valid RFC3339 timestamp") {
		return false
	}

	// Calculate the absolute difference
	diff := expectedTime.Sub(actualTime)
	if diff < 0 {
		diff = -diff
	}

	// Assert that the difference is GREATER than the tolerance
	return assert.Greater(t, diff, time.Millisecond, msgAndArgs...)
}
