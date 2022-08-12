package assert

import (
	"testing"

	"github.com/qw20012/go-basic"
)

// Assert that two objects are equal. Log and terminate when failed.
//
// Example:
// Equal(t, "DD", Repeat("D", 2))
// Equal(t, "DD", Repeat("D", 2), "%s expected Type=%v, Got=%v", "Test", "DD", Repeat("D", 2))
func Equal(t *testing.T, expected, actual any, msg ...any) {
	basic.HandleTest(t.Fatalf, expected, actual, true, msg)
}

// Assert that two objects are not equal. Log and terminate when failed.
//
// Example:
// NotEqual(t, "DD", Repeat("D", 2))
// NotEqual(t, "DD", Repeat("D", 2), "%s expected Type=%v, Got=%v", "Test", "DD", Repeat("D", 2))
func NotEqual(t *testing.T, expected, actual any, msg ...any) {
	basic.HandleTest(t.Fatalf, expected, actual, false, msg)
}

// Assert that given object is true. Log and terminate when failed.
//
// Example:
// True(t, CanRepeat())
// True(t, CanRepeat(), "%s expected Type=%v, Got=%v", "Test", true, CanRepeat())
func True(t *testing.T, actual bool, msg ...any) {
	basic.HandleTest(t.Fatalf, true, actual, true, msg)
}

// Assert that given object is false. Log and terminate when failed.
//
// Example:
// False(t, CanRepeat())
// False(t, CanRepeat(), "%s expected Type=%v, Got=%v", "Test", true, CanRepeat())
func False(t *testing.T, actual bool, msg ...any) {
	basic.HandleTest(t.Fatalf, false, actual, true, msg)
}
