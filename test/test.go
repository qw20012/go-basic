package test

import (
	"testing"

	"github.com/qw20012/go-basic"
)

// Test that two objects are equal. Log and continue when failed.
//
// Example:
// Equal(t, "DD", Repeat("D", 2))
// Equal(t, "DD", Repeat("D", 2), "%s expected Type=%v, Got=%v", "Test", "DD", Repeat("D", 2))
func Equal(t *testing.T, expected, actual any, msg ...any) bool {
	return basic.HandleTest(t.Errorf, expected, actual, true, msg)
}

// Test that two objects are not equal. Log and continue when failed.
//
// Example:
// NotEqual(t, "DD", Repeat("D", 2))
// NotEqual(t, "DD", Repeat("D", 2), "%s expected Type=%v, Got=%v", "Test", "DD", Repeat("D", 2))
func NotEqual(t *testing.T, expected, actual any, msg ...any) bool {
	return basic.HandleTest(t.Errorf, expected, actual, false, msg)
}

// Test that given object is true. Log and continue when failed.
//
// Example:
// True(t, CanRepeat())
// True(t, CanRepeat(), "%s expected Type=%v, Got=%v", "Test", true, CanRepeat())
func True(t *testing.T, actual bool, msg ...any) bool {
	return basic.HandleTest(t.Errorf, true, actual, true, msg)
}

// Test that given object is false. Log and continue when failed.
//
// Example:
// False(t, CanRepeat())
// False(t, CanRepeat(), "%s expected Type=%v, Got=%v", "Test", true, CanRepeat())
func False(t *testing.T, actual bool, msg ...any) bool {
	return basic.HandleTest(t.Errorf, false, actual, true, msg)
}
